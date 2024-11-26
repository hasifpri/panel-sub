package corerepository

import "github.com/bytesaddict/dancok"

type QueryInfo struct {
	Filter          string
	Sort            string
	SelectParameter dancok.SelectParameter
}
