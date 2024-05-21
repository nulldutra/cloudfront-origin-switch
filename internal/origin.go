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

func ListOrigins(cloudfrontID string) {
	session := NewSession()
	svc := cloudfront.New(session)

	input := &cloudfront.GetDistributionInput{
		Id: aws.String(cloudfrontID),
	}

	distributions, err := svc.GetDistribution(input)
	if err != nil {
		log.Panic(err)
	}

	for _, origin := range distributions.Distribution.DistributionConfig.Origins.Items {
		fmt.Printf("[+] Origin ID: %s - Domain Name: %s\n", *origin.Id, *origin.DomainName)
	}
}
