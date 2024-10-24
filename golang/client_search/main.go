package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	pb "distributed-system"

	"google.golang.org/grpc"
)

func saveImage(pet string, imageData []byte) (string, error) {
	// create image direction
	imageDir := "downloaded_images"
	if _, err := os.Stat(imageDir); os.IsNotExist(err) {
		os.Mkdir(imageDir, os.ModePerm)
	}

	// save image
	imagePath := filepath.Join(imageDir, pet+".jpg")
	err := os.WriteFile(imagePath, imageData, 0644)
	if err != nil {
		return "", fmt.Errorf("unable to save picture: %v", err)
	}

	fmt.Printf("save image: %s\n", imagePath)
	return imagePath, nil
}

func openImage(imagePath string) error {
	// open picture command according to os
	var cmd *exec.Cmd
	switch os := runtime.GOOS; os {
	case "darwin": // macOS
		cmd = exec.Command("open", imagePath)
	case "linux": // Linux
		cmd = exec.Command("xdg-open", imagePath)
	case "windows": // Windows
		cmd = exec.Command("cmd", "/c", "start", imagePath)
	default:
		return fmt.Errorf("unsupport os: %s", os)
	}

	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("open picture error: %v", err)
	}

	return nil
}

func searchPet(client pb.PetServiceClient) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("please input search keyword like: name:Labrador or gender:Female:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// parse input
	splitInput := strings.Split(input, ":")
	if len(splitInput) != 2 {
		log.Fatalf("invalid type, please use this format, for example: name:Labrador")
	}
	field := splitInput[0]
	keyword := splitInput[1]

	var searchRequest *pb.SearchPetRequest

	// 根据输入的关键字构造对应的搜索请求
	switch field {
	case "name":
		searchRequest = &pb.SearchPetRequest{
			Detail: &pb.SearchPetRequest_Name{
				Name: keyword,
			},
		}
	case "gender":
		searchRequest = &pb.SearchPetRequest{
			Detail: &pb.SearchPetRequest_Gender{
				Gender: keyword,
			},
		}
	case "age":
		age, err := strconv.Atoi(keyword)
		if err != nil {
			log.Fatalf("无效的年龄: %v", err)
		}
		searchRequest = &pb.SearchPetRequest{
			Detail: &pb.SearchPetRequest_Age{
				Age: int32(age),
			},
		}
	case "breed":
		searchRequest = &pb.SearchPetRequest{
			Detail: &pb.SearchPetRequest_Breed{
				Breed: keyword,
			},
		}
	default:
		log.Fatalf("invalid field: %s", field)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// call searchPet function
	resp, err := client.SearchPet(ctx, searchRequest)
	if err != nil {
		log.Fatalf("unable to search: %v", err)
	}

	// print result
	log.Printf("find %d pets:", len(resp.Pets))
	for _, pet := range resp.Pets {
		log.Printf("name: %s, gender: %s, age: %d, breed: %s", pet.Name, pet.Gender, pet.Age, pet.Breed)
		imagePath, err := saveImage(fmt.Sprintf("%s_%s_%d_%s", pet.Name, pet.Gender, pet.Age, pet.Breed), pet.Picture)
		if err != nil {
			log.Printf("save picture error: %v", err)
		}

		// open image
		err = openImage(imagePath)
		if err != nil {
			log.Printf("open picture error: %v", err)
		}
	}
}

func main() {
	// Establish a gRPC connection to the 'server-container' at port 50051.
	// server-container is my server's container name
	conn, err := grpc.Dial("server-container:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
	}
	defer conn.Close()

	// Create a new gRPC client
	client := pb.NewPetServiceClient(conn)

	searchPet(client)
}
