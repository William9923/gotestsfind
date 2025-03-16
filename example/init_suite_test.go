package example_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ExampleTestSuite struct {
	suite.Suite
	VariableThatShouldStartAtFive int
}

func (suite *ExampleTestSuite) SetupTest() {
	suite.VariableThatShouldStartAtFive = 5
}

func (suite *ExampleTestSuite) TestExampleFailed() {
	assert.Equal(suite.T(), 10, suite.VariableThatShouldStartAtFive)
	suite.Equal(10, suite.VariableThatShouldStartAtFive)
}

func (suite *ExampleTestSuite) TestExample() {
	assert.Equal(suite.T(), 5, suite.VariableThatShouldStartAtFive)
	suite.Equal(5, suite.VariableThatShouldStartAtFive)
}

func (suite *ExampleTestSuite) TestTableDrivenSubtests() {
	testCases := []struct {
		name           string
		input          int
		expectedOutput int
	}{
		{
			name:           "Input is 0",
			input:          0,
			expectedOutput: 0,
		},
		{
			name:           "Input is 1",
			input:          1,
			expectedOutput: 1,
		},
		{
			name:           "Input is 5",
			input:          5,
			expectedOutput: 5,
		},
		{
			name:           "Input is 10",
			input:          10,
			expectedOutput: 15,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.VariableThatShouldStartAtFive = tc.input
			suite.Equal(tc.expectedOutput, suite.VariableThatShouldStartAtFive)
		})
	}
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}
