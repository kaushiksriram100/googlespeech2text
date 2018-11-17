//Speech to Text testing with a service account

package main

import (
	speech "cloud.google.com/go/speech/apiv1"
	"context"
	"flag"
	"fmt"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
	"io/ioutil"
	"log"
)

func main() {

	//get the input file.

	var speech_file *string = flag.String("input", "/tmp/GCP_AI/somespeech.raw", "the input *.raw file. To convert - sox bob1.wav --channels=1 --rate 16k --bits 16 bob1a.raw")
	flag.Parse()
	//creates a context that never expires
	ctx := context.Background()

	//create a client with the above context
	client, err := speech.NewClient(ctx)

	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	// Set the name of the audio file to transcribe

	//Read the audio file into the memory. ReadFile return []bytes (go style!)
	data, err := ioutil.ReadFile(*speech_file)

	if err != nil {
		log.Fatalf("failed to read audio file. May be corrupted. %v", err)
	}

	//Create a speech config

	speech_conf := &speechpb.RecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			Encoding:        speechpb.RecognitionConfig_LINEAR16,
			SampleRateHertz: 16000,
			LanguageCode:    "en-US",
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Content{Content: data},
		},
	}
	//detect speech portions of the audio file
	resp, err := client.Recognize(ctx, speech_conf)

	if err != nil {
		log.Fatalf("failed to recognize: %v", err)

	}

	for _, result := range resp.Results {
		for _, alt := range result.Alternatives {
			fmt.Printf("\"%v\" (confidence=%3f)\n", alt.Transcript, alt.Confidence)
		}
	}
	return

}
