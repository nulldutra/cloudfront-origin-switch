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

// switchCmd represents the switch command
var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Switch origins on CloudFront distribution",
	Run: func(cmd *cobra.Command, args []string) {
		internal.Switch(cloudFrontID, fromOriginID, toOriginId)
	},
}

func init() {
	rootCmd.AddCommand(switchCmd)
	switchCmd.Flags().StringVarP(&cloudFrontID, "id", "i", "", "The CloudFront ID")
	switchCmd.Flags().StringVarP(&fromOriginID, "from", "f", "", "The CloudFront current origin ID on distribution")
	switchCmd.Flags().StringVarP(&toOriginId, "to", "t", "", "The CloudFront next origin ID to switch on distribution")

	switchCmd.MarkFlagRequired("id")
	switchCmd.MarkFlagRequired("to")
}
