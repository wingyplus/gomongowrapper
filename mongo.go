// Copyright 2018, OpenCensus Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mongowrapper

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(ctx context.Context, opts ...*options.ClientOptions) (*WrappedClient, error) {
	ctx, span := roundtripTrackingSpan(ctx, "go.mongodb.org/mongo-driver.Connect")
	defer span.end(ctx)

	cc, err := mongo.NewClient(opts...)
	if err != nil {
		span.setError(ctx, err)
		return nil, err
	}

	wc := &WrappedClient{cc: cc}
	if err := wc.Connect(ctx); err != nil {
		span.setError(ctx, err)
	}

	return wc, err
}
