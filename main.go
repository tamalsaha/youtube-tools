// Sample Go code for user authorization

package main

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/context"
	_ "gomodules.xyz/gdrive-utils"
	gdrive_utils "gomodules.xyz/gdrive-utils"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func handleError(err error, message string) {
	if message == "" {
		message = "Error making API call"
	}
	if err != nil {
		log.Fatalf(message+": %v", err.Error())
	}
}

// https://developers.google.com/youtube/v3/guides/working_with_channel_ids
// Use channel id
// Youtube Studio > Customization > Basic Info
func channelsListByUsername(service *youtube.Service, parts []string, forUsername string) {
	call := service.Channels.List(parts)
	call = call.Id(forUsername)
	response, err := call.Do()
	handleError(err, "")

	fmt.Println(fmt.Sprintf("This channel's ID is %s. Its title is '%s', "+
		"and it has %d views.",
		response.Items[0].Id,
		response.Items[0].Snippet.Title,
		response.Items[0].Statistics.ViewCount))
}

func main() {
	ctx := context.Background()

	client, err := gdrive_utils.DefaultClient(".", youtube.YoutubeReadonlyScope)
	handleError(err, "Error creating YouTube client")
	service, err := youtube.NewService(ctx, option.WithHTTPClient(client))
	handleError(err, "Error creating YouTube client")

	channelsListByUsername(service, strings.Split("snippet,contentDetails,statistics", ","), "UCxObRDZ0DtaQe_cCP-dN-xg")
}
