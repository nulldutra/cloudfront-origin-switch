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

package cmd

import (
	"cloudfront-origin-switch/internal"

	"github.com/spf13/cobra"
)

// originsCmd represents the origins command
var originsCmd = &cobra.Command{
	Use:   "origins",
	Short: "List all origins in a CloudFront distribution",
	Run: func(cmd *cobra.Command, args []string) {
		internal.ListOrigins(cloudFrontID)
	},
}

func init() {
	listCmd.AddCommand(originsCmd)
	originsCmd.Flags().StringVarP(&cloudFrontID, "id", "i", "", "The CloudFront ID")
	originsCmd.MarkFlagRequired("id")
}
