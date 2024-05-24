package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/riweston/gotempl-cv/pkgs/types"
	"gopkg.in/yaml.v3"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("User ID argument is required")
	}

	url := fmt.Sprintf("https://api.credly.com/users/%s/badges", os.Args[1])
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error performing request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error: %s", resp.Status)
	}

	var credlyData types.CredlyData
	if err := json.NewDecoder(resp.Body).Decode(&credlyData); err != nil {
		log.Fatalf("Error decoding response: %v", err)
	}

	saveBadgeImages(credlyData)
	saveCertificates(credlyData)
}

func saveBadgeImages(credlyData types.CredlyData) {
	for _, data := range credlyData.Data {
		imagePath := filepath.Join("public", "images", "certificate-badges", fmt.Sprintf("%s.png", data.BadgeTemplate.VanitySlug))

		if _, err := os.Stat(imagePath); err == nil {
			log.Printf("Badge image already saved: %s", data.BadgeTemplate.Name)
			continue
		}

		if err := downloadFile(data.BadgeTemplate.ImageUrl, imagePath); err != nil {
			log.Fatalf("Error downloading badge image: %v", err)
		}

		log.Printf("Badge image saved: %s", data.BadgeTemplate.VanitySlug)
	}
}

func downloadFile(url, path string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error downloading file: %w", err)
	}
	defer resp.Body.Close()

	out, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		return fmt.Errorf("error saving file: %w", err)
	}

	return nil
}

func saveCertificates(credlyData types.CredlyData) {
	filePath := filepath.Join("public", "data", "certificates_gen.yml")
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Error creating YAML file: %v", err)
	}
	defer file.Close()

	var certificateData types.CertificateData
	for _, data := range credlyData.Data {
		certificateData.Certificates = append(certificateData.Certificates, types.Certificate{
			DisplayName: data.BadgeTemplate.Name,
			Issuer:      data.Issuer.Entities[0].Entity.Name,
			BadgePath:   filepath.Join("public", "images", "certificate-badges", fmt.Sprintf("%s.png", data.BadgeTemplate.VanitySlug)),
			IssueDate:   data.IssuedAt.Format("2006-01-02"),
			ExpireDate:  data.ExpiresAt.Format("2006-01-02"),
		})
	}

	if err := yaml.NewEncoder(file).Encode(&certificateData); err != nil {
		log.Fatalf("Error encoding YAML: %v", err)
	}

	log.Printf("Certificates saved to %s", filePath)
}
