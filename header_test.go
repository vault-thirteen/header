////////////////////////////////////////////////////////////////////////////////
//
// Copyright © 2019 by Vault Thirteen.
//
// All rights reserved. No part of this publication may be reproduced,
// distributed, or transmitted in any form or by any means, including
// photocopying, recording, or other electronic or mechanical methods,
// without the prior written permission of the publisher, except in the case
// of brief quotations embodied in critical reviews and certain other
// noncommercial uses permitted by copyright law. For permission requests,
// write to the publisher, addressed “Copyright Protected Material” at the
// address below.
//
////////////////////////////////////////////////////////////////////////////////
//
// Web Site Address:	https://github.com/vault-thirteen.
//
////////////////////////////////////////////////////////////////////////////////

//+build test

package header

import (
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_MakeListOfHeaders(t *testing.T) {

	var aTest *tester.Test
	var headers = []string{"aa", "bb", "cc"}
	var result string

	aTest = tester.New(t)
	result = MakeListOfHeaders(headers)
	aTest.MustBeEqual(result, "aa, bb, cc")
}
