/*
Copyright © 2024 Gabriel M. Dutra <nulldutra@proton.me>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package internal

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudfront"
)

func Switch(cloudfrontID, fromOriginID, toOriginId string) {
	session := NewSession()
	svc := cloudfront.New(session)

	result, err := svc.GetDistributionConfig(&cloudfront.GetDistributionConfigInput{
		Id: aws.String(cloudfrontID),
	})
	if err != nil {
		log.Fatalf("Failed to get distribution config: %v", err)
	}

	newOriginID := toOriginId

	for _, cacheBehavior := range result.DistributionConfig.CacheBehaviors.Items {
		if *cacheBehavior.TargetOriginId == fromOriginID {
			cacheBehavior.TargetOriginId = aws.String(newOriginID)
		}
	}

	if *result.DistributionConfig.DefaultCacheBehavior.TargetOriginId == fromOriginID {
		result.DistributionConfig.DefaultCacheBehavior.TargetOriginId = aws.String(newOriginID)
	}

	_, err = svc.UpdateDistribution(&cloudfront.UpdateDistributionInput{
		Id:                 aws.String(cloudfrontID),
		IfMatch:            result.ETag,
		DistributionConfig: result.DistributionConfig,
	})
	if err != nil {
		log.Fatalf("Failed to update distribution: %v", err)
	}

	fmt.Printf("[+] switching origin on CloudFront (%s) to %s\n", cloudfrontID, newOriginID)
}
