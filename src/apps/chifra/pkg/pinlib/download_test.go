// Copyright 2021 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.

package pinlib

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var manifestJSONSource = `
{
  "version": "2",
  "chain": "mainnet",
  "chainId": 1,
  "indexFormat": "Qmart6XP9XjL43p72PGR93QKytbK8jWWcMguhFgxATTya2",
  "bloomFormat": "QmNhPk39DUFoEdhUmtGARqiFECUHeghyeryxZM9kyRxzHD",
  "commitHash": "f29699d3281e41cb011ddfbe50b7f01bfe5e3c53",
  "names": "QmP4i6ihnVrj8Tx7cTFw4aY6ungpaPYxDJEZ7Vg1RSNSdm",
  "timestamps": "QmcvjroTiE95LWeiP8HHq1YA3ysRchLuVx8HLQui8WcSBV",
  "blockRange": "000000000-000864336",
  "pins": [
    {
      "fileName": "000000000-000000000",
      "bloomHash": "QmPQEgUm7nzQuW9HYyWp5Ff3aoUwg2rsxDngyuyddJTvrv",
      "indexHash": "QmZ5Atm8Z7aFLz2EycK4pVMHuH4y3PDGspuFejnE9fx2i5"
    },
    {
      "fileName": "000000001-000350277",
      "bloomHash": "QmZgrWAJLidkHJRLDVoZGCWAgmmcQEDCDM65XL5ZbAXxCM",
      "indexHash": "QmP1KvDPUJ1MqsCYcicJgTf5sxN7WjT7dZsrfBk2Jg3mSe"
    }
   ]
}
`

func TestDownloadJSON(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, manifestJSONSource)
	}))

	defer ts.Close()

	manifest, err := DownloadManifest(ts.URL)
	if err != nil {
		t.Error(err)
	}

	if l := len(manifest.Pins); l != 2 {
		t.Errorf("Wrong Pins length: %d", l)
	}
}
