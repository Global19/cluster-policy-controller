// Copyright ©2014 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package distuv

const (
	// oneOverRoot2Pi is the value of 1/(2Pi)^(1/2)
	// http://www.wolframalpha.com/input/?i=1%2F%282+*+pi%29%5E%281%2F2%29
	oneOverRoot2Pi = 0.39894228040143267793994605993438186847585863116493465766592582967065792589930183850125233390730693643030255886263518268

	//LogRoot2Pi is the value of log(sqrt(2*Pi))
	logRoot2Pi    = 0.91893853320467274178032973640561763986139747363778341281715154048276569592726039769474329863595419762200564662463433744
	negLogRoot2Pi = -logRoot2Pi
	log2Pi        = 1.8378770664093454835606594728112352797227949472755668
	ln2           = 0.69314718055994530941723212145817656807550013436025525412068000949339362196969471560586332699641868754200148102057068573368552023

	// Euler–Mascheroni constant.
	eulerGamma = 0.5772156649015328606065120900824024310421593359399235988057672348848677267776646709369470632917467495146314472498070824809605
)

const (
	panicNameMismatch = "parameter name mismatch"
)
