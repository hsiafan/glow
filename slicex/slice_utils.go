package slicex

import (
	"fmt"
	"reflect"
)

// Remove element at index
func RemoveStringAt(ptr *[]string, index int) error {
	slice := *ptr
	l := len(slice)
	if index < 0 || index >= l {
		return fmt.Errorf("invalid index: %v", index)
	}
	copy(slice[index:l-1], slice[index+1:l])
	*ptr = slice[:l-1]
	return nil
}

// Remove element at index
func RemoveBoolAt(ptr *[]bool, index int) error {
	slice := *ptr
	l := len(slice)
	if index < 0 || index >= l {
		return fmt.Errorf("invalid index: %v", index)
	}
	copy(slice[index:l-1], slice[index+1:l])
	*ptr = slice[:l-1]
	return nil
}

// Remove element at index
func RemoveIntAt(ptr *[]int, index int) error {
	slice := *ptr
	l := len(slice)
	if index < 0 || index >= l {
		return fmt.Errorf("invalid index: %v", index)
	}
	copy(slice[index:l-1], slice[index+1:l])
	*ptr = slice[:l-1]
	return nil
}

// Remove element at index
func RemoveInt64At(ptr *[]int64, index int) error {
	slice := *ptr
	l := len(slice)
	if index < 0 || index >= l {
		return fmt.Errorf("invalid index: %v", index)
	}
	copy(slice[index:l-1], slice[index+1:l])
	*ptr = slice[:l-1]
	return nil
}

// Remove element at index
func RemoveInt32At(ptr *[]int32, index int) error {
	slice := *ptr
	l := len(slice)
	if index < 0 || index >= l {
		return fmt.Errorf("invalid index: %v", index)
	}
	copy(slice[index:l-1], slice[index+1:l])
	*ptr = slice[:l-1]
	return nil
}

// Remove element at index
func RemoveUintAt(ptr *[]uint, index int) error {
	slice := *ptr
	l := len(slice)
	if index < 0 || index >= l {
		return fmt.Errorf("invalid index: %v", index)
	}
	copy(slice[index:l-1], slice[index+1:l])
	*ptr = slice[:l-1]
	return nil
}

// Remove element at index
func RemoveUint64At(ptr *[]uint64, index int) error {
	slice := *ptr
	l := len(slice)
	if index < 0 || index >= l {
		return fmt.Errorf("invalid index: %v", index)
	}
	copy(slice[index:l-1], slice[index+1:l])
	*ptr = slice[:l-1]
	return nil
}

// Remove element at index
func RemoveUint32At(ptr *[]uint32, index int) error {
	slice := *ptr
	l := len(slice)
	if index < 0 || index >= l {
		return fmt.Errorf("invalid index: %v", index)
	}
	copy(slice[index:l-1], slice[index+1:l])
	*ptr = slice[:l-1]
	return nil
}

// Remove element at index
func RemoveFloat64At(ptr *[]float64, index int) error {
	slice := *ptr
	l := len(slice)
	if index < 0 || index >= l {
		return fmt.Errorf("invalid index: %v", index)
	}
	copy(slice[index:l-1], slice[index+1:l])
	*ptr = slice[:l-1]
	return nil
}

// Remove element at index
func RemoveFloat32At(ptr *[]float32, index int) error {
	slice := *ptr
	l := len(slice)
	if index < 0 || index >= l {
		return fmt.Errorf("invalid index: %v", index)
	}
	copy(slice[index:l-1], slice[index+1:l])
	*ptr = slice[:l-1]
	return nil
}

// Remove element at index
func RemoveRuneAt(ptr *[]rune, index int) error {
	slice := *ptr
	l := len(slice)
	if index < 0 || index >= l {
		return fmt.Errorf("invalid index: %v", index)
	}
	copy(slice[index:l-1], slice[index+1:l])
	*ptr = slice[:l-1]
	return nil
}

// Remove element at index
func RemoveAt(ptr interface{}, index int) error {
	rv := reflect.ValueOf(ptr)
	slice := rv.Elem()
	l := slice.Len()
	if index < 0 || index >= l {
		return fmt.Errorf("invalid index: %v", index)
	}
	reflect.Copy(slice.Slice(index, l-1), slice.Slice(index+1, l))
	rv.Elem().Set(slice.Slice(0, l-1))
	return nil
}

// Remove element at index
// Remove all elements match the given condition
func RemoveStringIf(ptr *[]string, f func(idx int, v string) bool) {
	slice := *ptr
	var target = 0
	for idx, v := range slice {
		if !f(idx, v) {
			slice[target] = v
			target++
		}
	}
	if target != len(slice) {
		*ptr = slice[:target]
	}
}

// Remove element at index
// Remove all elements match the given condition
func RemoveBoolIf(ptr *[]bool, f func(idx int, v bool) bool) {
	slice := *ptr
	var target = 0
	for idx, v := range slice {
		if !f(idx, v) {
			slice[target] = v
			target++
		}
	}
	if target != len(slice) {
		*ptr = slice[:target]
	}
}

// Remove element at index
// Remove all elements match the given condition
func RemoveIntIf(ptr *[]int, f func(idx int, v int) bool) {
	slice := *ptr
	var target = 0
	for idx, v := range slice {
		if !f(idx, v) {
			slice[target] = v
			target++
		}
	}
	if target != len(slice) {
		*ptr = slice[:target]
	}
}

// Remove element at index
// Remove all elements match the given condition
func RemoveInt64If(ptr *[]int64, f func(idx int, v int64) bool) {
	slice := *ptr
	var target = 0
	for idx, v := range slice {
		if !f(idx, v) {
			slice[target] = v
			target++
		}
	}
	if target != len(slice) {
		*ptr = slice[:target]
	}
}

// Remove element at index
// Remove all elements match the given condition
func RemoveInt32If(ptr *[]int32, f func(idx int, v int32) bool) {
	slice := *ptr
	var target = 0
	for idx, v := range slice {
		if !f(idx, v) {
			slice[target] = v
			target++
		}
	}
	if target != len(slice) {
		*ptr = slice[:target]
	}
}

// Remove element at index
// Remove all elements match the given condition
func RemoveUintIf(ptr *[]uint, f func(idx int, v uint) bool) {
	slice := *ptr
	var target = 0
	for idx, v := range slice {
		if !f(idx, v) {
			slice[target] = v
			target++
		}
	}
	if target != len(slice) {
		*ptr = slice[:target]
	}
}

// Remove element at index
// Remove all elements match the given condition
func RemoveUint64If(ptr *[]uint64, f func(idx int, v uint64) bool) {
	slice := *ptr
	var target = 0
	for idx, v := range slice {
		if !f(idx, v) {
			slice[target] = v
			target++
		}
	}
	if target != len(slice) {
		*ptr = slice[:target]
	}
}

// Remove element at index
// Remove all elements match the given condition
func RemoveUint32If(ptr *[]uint32, f func(idx int, v uint32) bool) {
	slice := *ptr
	var target = 0
	for idx, v := range slice {
		if !f(idx, v) {
			slice[target] = v
			target++
		}
	}
	if target != len(slice) {
		*ptr = slice[:target]
	}
}

// Remove element at index
// Remove all elements match the given condition
func RemoveFloat64If(ptr *[]float64, f func(idx int, v float64) bool) {
	slice := *ptr
	var target = 0
	for idx, v := range slice {
		if !f(idx, v) {
			slice[target] = v
			target++
		}
	}
	if target != len(slice) {
		*ptr = slice[:target]
	}
}

// Remove element at index
// Remove all elements match the given condition
func RemoveFloat32If(ptr *[]float32, f func(idx int, v float32) bool) {
	slice := *ptr
	var target = 0
	for idx, v := range slice {
		if !f(idx, v) {
			slice[target] = v
			target++
		}
	}
	if target != len(slice) {
		*ptr = slice[:target]
	}
}

// Remove element at index
// Remove all elements match the given condition
func RemoveRuneIf(ptr *[]rune, f func(idx int, v rune) bool) {
	slice := *ptr
	var target = 0
	for idx, v := range slice {
		if !f(idx, v) {
			slice[target] = v
			target++
		}
	}
	if target != len(slice) {
		*ptr = slice[:target]
	}
}

func RemoveIf(ptr interface{}, f func(idx int, v interface{}) bool) {
	rv := reflect.ValueOf(ptr)
	slice := rv.Elem()
	var target = 0
	for idx := 0; idx < slice.Len(); idx++ {
		iv := slice.Index(idx)
		if !f(idx, iv.Interface()) {
			slice.Index(target).Set(iv)
			target++
		}
	}
	if target != slice.Len() {
		rv.Elem().Set(slice.Slice(0, target))
	}
}

// Find first appearance of value. If not found, return -1
func FindString(slice []string, v string) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := 0; i < len(slice); i++ {
		if slice[i] == v {
			return i
		}
	}
	return -1
}

// Find first appearance of value. If not found, return -1
func FindBool(slice []bool, v bool) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := 0; i < len(slice); i++ {
		if slice[i] == v {
			return i
		}
	}
	return -1
}

// Find first appearance of value. If not found, return -1
func FindInt(slice []int, v int) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := 0; i < len(slice); i++ {
		if slice[i] == v {
			return i
		}
	}
	return -1
}

// Find first appearance of value. If not found, return -1
func FindInt64(slice []int64, v int64) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := 0; i < len(slice); i++ {
		if slice[i] == v {
			return i
		}
	}
	return -1
}

// Find first appearance of value. If not found, return -1
func FindInt32(slice []int32, v int32) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := 0; i < len(slice); i++ {
		if slice[i] == v {
			return i
		}
	}
	return -1
}

// Find first appearance of value. If not found, return -1
func FindUint(slice []uint, v uint) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := 0; i < len(slice); i++ {
		if slice[i] == v {
			return i
		}
	}
	return -1
}

// Find first appearance of value. If not found, return -1
func FindUint64(slice []uint64, v uint64) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := 0; i < len(slice); i++ {
		if slice[i] == v {
			return i
		}
	}
	return -1
}

// Find first appearance of value. If not found, return -1
func FindUint32(slice []uint32, v uint32) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := 0; i < len(slice); i++ {
		if slice[i] == v {
			return i
		}
	}
	return -1
}

// Find first appearance of value. If not found, return -1
func FindFloat64(slice []float64, v float64) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := 0; i < len(slice); i++ {
		if slice[i] == v {
			return i
		}
	}
	return -1
}

// Find first appearance of value. If not found, return -1
func FindFloat32(slice []float32, v float32) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := 0; i < len(slice); i++ {
		if slice[i] == v {
			return i
		}
	}
	return -1
}

// Find first appearance of value. If not found, return -1
func FindRune(slice []rune, v rune) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := 0; i < len(slice); i++ {
		if slice[i] == v {
			return i
		}
	}
	return -1
}

// Find first appearance of value. If not found, return -1
func Find(slice interface{}, v interface{}) int {
	sc := reflect.ValueOf(slice)
	if !sc.IsValid() || sc.IsZero() || sc.IsNil() || sc.Len() == 0 {
		return -1
	}
	l := sc.Len()
	for i := 0; i < l; i++ {
		if sc.Index(i).Interface() == v {
			return i
		}
	}
	return -1
}

// Find next appearance of value, from end to beginning. If not found, return -1
func ReverseFindString(slice []string, v string) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == v {
			return i
		}
	}
	return -1
}

// Find next appearance of value, from end to beginning. If not found, return -1
func ReverseFindBool(slice []bool, v bool) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == v {
			return i
		}
	}
	return -1
}

// Find next appearance of value, from end to beginning. If not found, return -1
func ReverseFindInt(slice []int, v int) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == v {
			return i
		}
	}
	return -1
}

// Find next appearance of value, from end to beginning. If not found, return -1
func ReverseFindInt64(slice []int64, v int64) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == v {
			return i
		}
	}
	return -1
}

// Find next appearance of value, from end to beginning. If not found, return -1
func ReverseFindInt32(slice []int32, v int32) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == v {
			return i
		}
	}
	return -1
}

// Find next appearance of value, from end to beginning. If not found, return -1
func ReverseFindUint(slice []uint, v uint) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == v {
			return i
		}
	}
	return -1
}

// Find next appearance of value, from end to beginning. If not found, return -1
func ReverseFindUint64(slice []uint64, v uint64) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == v {
			return i
		}
	}
	return -1
}

// Find next appearance of value, from end to beginning. If not found, return -1
func ReverseFindUint32(slice []uint32, v uint32) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == v {
			return i
		}
	}
	return -1
}

// Find next appearance of value, from end to beginning. If not found, return -1
func ReverseFindFloat64(slice []float64, v float64) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == v {
			return i
		}
	}
	return -1
}

// Find next appearance of value, from end to beginning. If not found, return -1
func ReverseFindFloat32(slice []float32, v float32) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == v {
			return i
		}
	}
	return -1
}

// Find next appearance of value, from end to beginning. If not found, return -1
func ReverseFindRune(slice []rune, v rune) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == v {
			return i
		}
	}
	return -1
}

// Find next appearance of value, from end to beginning. If not found, return -1
func ReverseFind(slice interface{}, v interface{}) int {
	sc := reflect.ValueOf(slice)
	if !sc.IsValid() || sc.IsZero() || sc.IsNil() || sc.Len() == 0 {
		return -1
	}
	l := sc.Len()
	for i := l - 1; i >= 0; i-- {
		if sc.Index(i).Interface() == v {
			return i
		}
	}
	return -1
}

// Find next appearance of value matching the condition, from end to beginning. If not found, return -1
func FindStringBy(slice []string, f func(v string) bool) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := 0; i <= len(slice); i++ {
		if f(slice[i]) {
			return i
		}
	}
	return -1
}

// Find next appearance of value matching the condition, from end to beginning. If not found, return -1
func FindBoolBy(slice []bool, f func(v bool) bool) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := 0; i <= len(slice); i++ {
		if f(slice[i]) {
			return i
		}
	}
	return -1
}

// Find next appearance of value matching the condition, from end to beginning. If not found, return -1
func FindIntBy(slice []int, f func(v int) bool) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := 0; i <= len(slice); i++ {
		if f(slice[i]) {
			return i
		}
	}
	return -1
}

// Find next appearance of value matching the condition, from end to beginning. If not found, return -1
func FindInt64By(slice []int64, f func(v int64) bool) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := 0; i <= len(slice); i++ {
		if f(slice[i]) {
			return i
		}
	}
	return -1
}

// Find next appearance of value matching the condition, from end to beginning. If not found, return -1
func FindInt32By(slice []int32, f func(v int32) bool) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := 0; i <= len(slice); i++ {
		if f(slice[i]) {
			return i
		}
	}
	return -1
}

// Find next appearance of value matching the condition, from end to beginning. If not found, return -1
func FindUintBy(slice []uint, f func(v uint) bool) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := 0; i <= len(slice); i++ {
		if f(slice[i]) {
			return i
		}
	}
	return -1
}

// Find next appearance of value matching the condition, from end to beginning. If not found, return -1
func FindUint64By(slice []uint64, f func(v uint64) bool) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := 0; i <= len(slice); i++ {
		if f(slice[i]) {
			return i
		}
	}
	return -1
}

// Find next appearance of value matching the condition, from end to beginning. If not found, return -1
func FindUint32By(slice []uint32, f func(v uint32) bool) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := 0; i <= len(slice); i++ {
		if f(slice[i]) {
			return i
		}
	}
	return -1
}

// Find next appearance of value matching the condition, from end to beginning. If not found, return -1
func FindFloat64By(slice []float64, f func(v float64) bool) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := 0; i <= len(slice); i++ {
		if f(slice[i]) {
			return i
		}
	}
	return -1
}

// Find next appearance of value matching the condition, from end to beginning. If not found, return -1
func FindFloat32By(slice []float32, f func(v float32) bool) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := 0; i <= len(slice); i++ {
		if f(slice[i]) {
			return i
		}
	}
	return -1
}

// Find next appearance of value matching the condition, from end to beginning. If not found, return -1
func FindRuneBy(slice []rune, f func(v rune) bool) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := 0; i <= len(slice); i++ {
		if f(slice[i]) {
			return i
		}
	}
	return -1
}

// Find next appearance of value matching the condition, from end to beginning. If not found, return -1
func FindBy(slice interface{}, f func(v interface{}) bool) int {
	sc := reflect.ValueOf(slice)
	if !sc.IsValid() || sc.IsZero() || sc.IsNil() || sc.Len() == 0 {
		return -1
	}
	l := sc.Len()
	for i := 0; i < l; i++ {
		if f(sc.Index(i).Interface()) {
			return i
		}
	}
	return -1
}

// Find next appearance of value matching the condition, from end to beginning. If not found, return -1
func ReverseFindStringBy(slice []string, f func(v string) bool) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := len(slice) - 1; i >= 0; i-- {
		if f(slice[i]) {
			return i
		}
	}
	return -1
}

// Find next appearance of value matching the condition, from end to beginning. If not found, return -1
func ReverseFindBoolBy(slice []bool, f func(v bool) bool) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := len(slice) - 1; i >= 0; i-- {
		if f(slice[i]) {
			return i
		}
	}
	return -1
}

// Find next appearance of value matching the condition, from end to beginning. If not found, return -1
func ReverseFindIntBy(slice []int, f func(v int) bool) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := len(slice) - 1; i >= 0; i-- {
		if f(slice[i]) {
			return i
		}
	}
	return -1
}

// Find next appearance of value matching the condition, from end to beginning. If not found, return -1
func ReverseFindInt64By(slice []int64, f func(v int64) bool) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := len(slice) - 1; i >= 0; i-- {
		if f(slice[i]) {
			return i
		}
	}
	return -1
}

// Find next appearance of value matching the condition, from end to beginning. If not found, return -1
func ReverseFindInt32By(slice []int32, f func(v int32) bool) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := len(slice) - 1; i >= 0; i-- {
		if f(slice[i]) {
			return i
		}
	}
	return -1
}

// Find next appearance of value matching the condition, from end to beginning. If not found, return -1
func ReverseFindUintBy(slice []uint, f func(v uint) bool) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := len(slice) - 1; i >= 0; i-- {
		if f(slice[i]) {
			return i
		}
	}
	return -1
}

// Find next appearance of value matching the condition, from end to beginning. If not found, return -1
func ReverseFindUint64By(slice []uint64, f func(v uint64) bool) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := len(slice) - 1; i >= 0; i-- {
		if f(slice[i]) {
			return i
		}
	}
	return -1
}

// Find next appearance of value matching the condition, from end to beginning. If not found, return -1
func ReverseFindUint32By(slice []uint32, f func(v uint32) bool) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := len(slice) - 1; i >= 0; i-- {
		if f(slice[i]) {
			return i
		}
	}
	return -1
}

// Find next appearance of value matching the condition, from end to beginning. If not found, return -1
func ReverseFindFloat64By(slice []float64, f func(v float64) bool) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := len(slice) - 1; i >= 0; i-- {
		if f(slice[i]) {
			return i
		}
	}
	return -1
}

// Find next appearance of value matching the condition, from end to beginning. If not found, return -1
func ReverseFindFloat32By(slice []float32, f func(v float32) bool) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := len(slice) - 1; i >= 0; i-- {
		if f(slice[i]) {
			return i
		}
	}
	return -1
}

// Find next appearance of value matching the condition, from end to beginning. If not found, return -1
func ReverseFindRuneBy(slice []rune, f func(v rune) bool) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	for i := len(slice) - 1; i >= 0; i-- {
		if f(slice[i]) {
			return i
		}
	}
	return -1
}

// Find next appearance of value matching the condition, from end to beginning. If not found, return -1
func ReverseFindBy(slice interface{}, f func(v interface{}) bool) int {
	sc := reflect.ValueOf(slice)
	if !sc.IsValid() || sc.IsZero() || sc.IsNil() || sc.Len() == 0 {
		return -1
	}
	l := sc.Len()
	for i := l - 1; i >= 0; i-- {
		if f(sc.Index(i).Interface()) {
			return i
		}
	}
	return -1
}
