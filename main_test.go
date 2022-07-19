package main
import (
	"testing"
	"github.com/stretchr/testify/assert"
) 

func TestPockerCaseOneTrue(t *testing.T){
	assert:= assert.New(t)
	numbers:=[]int{2, 3, 4 ,5, 6}
	assert.Equal(validatePockerHands(numbers), true)
}

func TestPockerCaseTwoTrue(t *testing.T){
	assert:= assert.New(t)
	numbers:=[]int{6, 5, 4 ,2, 3}
	assert.Equal(validatePockerHands(numbers), true)
}

func TestPockerCaseThreeFalse(t *testing.T){
	assert:= assert.New(t)
	numbers:=[]int{7, 7, 12 ,11, 3, 4, 14}
	assert.Equal(validatePockerHands(numbers), false)
}

func TestPockerCaseFourFalse(t *testing.T){
	assert:= assert.New(t)
	numbers:=[]int{7, 3, 2}
	assert.Equal(validatePockerHands(numbers), false)
}

func TestPockerCaseFiveFalse(t *testing.T){
	assert:= assert.New(t)
	numbers:=[]int{1}
	assert.Equal(validatePockerHands(numbers), false)
}

func TestPockerCaseSixFalse(t *testing.T){
	assert:= assert.New(t)
	numbers:=[]int{2,3,4,5,6,7,8,9}
	assert.Equal(validatePockerHands(numbers), false)
}

func TestPockerCaseSevenFalse(t *testing.T){
	assert:= assert.New(t)
	numbers:=[]int{9,10,11,13,14,15}
	assert.Equal(validatePockerHands(numbers), false)
}

func TestPockerCaseNineFalse(t *testing.T){
	assert:= assert.New(t)
	numbers:=[]int{1,2}
	assert.Equal(validatePockerHands(numbers), false)
}

func TestPockerCaseTenFalse(t *testing.T){
	assert:= assert.New(t)
	numbers:=[]int{}
	assert.Equal(validatePockerHands(numbers), false)
}

func TestPockerCaseElevenFalse(t *testing.T){
	assert:= assert.New(t)
	numbers:=[]int{2,3}
	assert.Equal(validatePockerHands(numbers), true)
}