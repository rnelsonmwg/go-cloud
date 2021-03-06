// Copyright 2018 The Go Cloud Development Kit Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package azureblob_test

import (
	"context"
	"log"

	"github.com/Azure/azure-storage-blob-go/azblob"
	"gocloud.dev/blob"
	"gocloud.dev/blob/azureblob"
)

func Example() {
	const (
		// Your Azure Storage Acount and Access Key.
		accountName = azureblob.AccountName("my-account")
		accountKey  = azureblob.AccountKey("my-account-key")
		// The storage container to access.
		containerName = "mycontainer"
	)

	// Create a credentials object.
	credential, err := azureblob.NewCredential(accountName, accountKey)
	if err != nil {
		log.Fatal(err)
	}

	// Create a Pipeline, using whatever PipelineOptions you need.
	pipeline := azureblob.NewPipeline(credential, azblob.PipelineOptions{})

	// Create a *blob.Bucket.
	// The credential Option is required if you're going to use blob.SignedURL.
	ctx := context.Background()
	b, err := azureblob.OpenBucket(ctx, pipeline, accountName, containerName, &azureblob.Options{Credential: credential})
	if err != nil {
		log.Fatal(err)
	}

	// Now we can use b to read or write files to the container.
	data, err := b.ReadAll(ctx, "my-key")
	if err != nil {
		log.Fatal(err)
	}
	_ = data
}

func Example_sasToken() {
	const (
		// Your Azure Storage Account and SASToken.
		accountName = azureblob.AccountName("my-account")
		sasToken    = azureblob.SASToken("my-SAS-token")
		// The storage container to access.
		containerName = "mycontainer"
	)

	// Since we're using a SASToken, we can use anonymous credentials.
	credential := azblob.NewAnonymousCredential()

	// Create a Pipeline, using whatever PipelineOptions you need.
	pipeline := azureblob.NewPipeline(credential, azblob.PipelineOptions{})

	// Create a *blob.Bucket.
	// Note that we're not supplying azureblob.Options.Credential, so SignedURL
	// won't work. To use SignedURL, you need a real credential (see the other
	// example).
	ctx := context.Background()
	b, err := azureblob.OpenBucket(ctx, pipeline, accountName, containerName, &azureblob.Options{SASToken: sasToken})
	if err != nil {
		log.Fatal(err)
	}

	// Now we can use b to read or write files to the container.
	data, err := b.ReadAll(ctx, "my-key")
	if err != nil {
		log.Fatal(err)
	}
	_ = data
}

func Example_open() {
	ctx := context.Background()

	// Open creates a *blob.Bucket from a URL.
	// This URL will open the container "mycontainer" using default
	// credentials found in the environment variables
	// AZURE_STORAGE_ACCOUNT plus at least one of AZURE_STORAGE_KEY
	// and AZURE_STORAGE_SAS_TOKEN.
	b, err := blob.OpenBucket(ctx, "azblob://mycontainer")
	_, _ = b, err
}
