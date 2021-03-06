package dbsearch

import (
	"reflect"
	"testing"
)

func Test_Array(t *testing.T) {
	s := init_test_data()
	if s != nil {
		_01_array_int(t, s)
		_02_array_int64(t, s)
		_03_array_uint64(t, s)
		_04_array_uint(t, s)
		_11_array_bool(t, s)
		_12_array_bool(t, s)
		_21_array_string(t, s)
		_31_array_float32(t, s)
		_32_array_float64(t, s)
	}
	//t.Fatal("Success [no error] test")
}

func array_main_f_test_table(s *Searcher) {
	sql_create := " CREATE TABLE public.test " +
		"(col1 int, col2 bigint[], col3 smallint[], col4 integer[], " +
		" col7 text[], col8 varchar(50)[], col9 char(10)[], " +
		" col11 real[], col12 double precision[], col13 numeric[], col14 decimal[], " +
		" col15 money[], col16 boolean[]  " +
		") "
	sql_cols := "INSERT INTO test(col1, col2, col3, col4, col7, col8, col9, col11, col12, col13, col14, col15, col16 ) "
	sql_vals := []string{
		" VALUES (1, '{9223372036854775807,3,2}'::bigint[],  '{18,28,33}'::smallint[],  '{884,-121}'::int[],  '{\"10\",\"Малененький текст\"}'::text[], " +
			" '{\" varchar \", \" varchar next\", \"654\"}','{\" char[] \",\" char2[] \",\"123.56234\"}', '{-12.13,13.13213409,-12.130909}', '{14.105,-15.015,-15.1500}', '{-16.17,16.1600,-16.173244}', " +
			" '{18.00,18.18,1800.180999234}'::decimal[], '{21.21,22.22,23.23}'::money[], '{TRUE,FALSE,TRUE,FALSE}'::boolean[] ) ",
		" VALUES (2, null, null, null, null, null, null, null, null, null, null, null, null ) ", // check null - nil
	}

	make_t_table(s, sql_create, sql_cols, sql_vals)
}

/*
	[]int test
*/
type array_int_TestPlace struct {
	Col1  int   `db:"col1" type:"int"`
	Col2  []int `db:"col2" type:"[]bigint"`
	Col3  []int `db:"col3" type:"[]smallint"`
	Col4  []int `db:"col4" type:"[]int"`
	Col7  []int `db:"col7" type:"[]text"`
	Col8  []int `db:"col8" type:"[]varchar"`
	Col9  []int `db:"col9" type:"[]char"`
	Col11 []int `db:"col11" type:"[]real"`
	Col12 []int `db:"col12" type:"[]double"`
	Col13 []int `db:"col13" type:"[]numeric"`
	Col14 []int `db:"col14" type:"[]decimal"`
	Col15 []int `db:"col15" type:"[]money"`
	Col16 []int `db:"col16" type:"[]bool"`
}

var array_int_mTestType *AllRows = &AllRows{
	SType: reflect.TypeOf(array_int_TestPlace{}),
}

func _01_array_int(t *testing.T, s *Searcher) {
	array_main_f_test_table(s)
	p := []array_int_TestPlace{}
	s.Get(array_int_mTestType, &p, "SELECT * FROM public.test ORDER BY 1")

	if p[0].Col2[0] != 9223372036854775807 {
		t.Fatal("Error array_int_TestPlace.Col2; []int <= []bigint")
	}
	if p[0].Col3[0] != 18 {
		t.Fatal("Error array_int_TestPlace.Col3; []int <= []smallint")
	}
	if p[0].Col4[0] != 884 {
		t.Fatal("Error array_int_TestPlace.Col4; []int <= []int")
	}
	if p[0].Col7[0] != 10 {
		t.Fatal("Error array_int_TestPlace.Col7; []int <= []text")
	}
	if p[0].Col8[1] != 0 {
		t.Fatal("Error array_int_TestPlace.Col8; []int <= []varchar")
	}
	if p[0].Col16[0] != 1 {
		t.Fatal("Error array_int_TestPlace.Col16; []int <= []bool")
	}
}

/*
	int64 test
*/
type array_int64_TestPlace struct {
	Col1  int     `db:"col1" type:"int"`
	Col2  []int64 `db:"col2" type:"[]bigint"`
	Col3  []int64 `db:"col3" type:"[]smallint"`
	Col4  []int64 `db:"col4" type:"[]int"`
	Col7  []int64 `db:"col7" type:"[]text"`
	Col8  []int64 `db:"col8" type:"[]varchar"`
	Col9  []int64 `db:"col9" type:"[]char"`
	Col11 []int64 `db:"col11" type:"[]real"`
	Col12 []int64 `db:"col12" type:"[]double"`
	Col13 []int64 `db:"col13" type:"[]numeric"`
	Col14 []int64 `db:"col14" type:"[]decimal"`
	Col15 []int64 `db:"col15" type:"[]money"`
	Col16 []int64 `db:"col16" type:"[]bool"`
}

var array_int64_mTestType *AllRows = &AllRows{
	SType: reflect.TypeOf(array_int64_TestPlace{}),
}

func _02_array_int64(t *testing.T, s *Searcher) {
	array_main_f_test_table(s)
	p := []array_int64_TestPlace{}
	s.Get(array_int64_mTestType, &p, "SELECT * FROM public.test ORDER BY 1")

	if p[0].Col2[0] != 9223372036854775807 {
		t.Fatal("Error array_int_TestPlace.Col2; []int <= []bigint")
	}
	if p[0].Col3[0] != 18 {
		t.Fatal("Error array_int_TestPlace.Col3; []int <= []smallint")
	}
	if p[0].Col4[0] != 884 {
		t.Fatal("Error array_int_TestPlace.Col4; []int <= []int")
	}
	if p[0].Col7[0] != 10 {
		t.Fatal("Error array_int_TestPlace.Col7; []int <= []text")
	}
	if p[0].Col8[1] != 0 {
		t.Fatal("Error array_int_TestPlace.Col8; []int <= []varchar")
	}
	if p[0].Col16[0] != 1 {
		t.Fatal("Error array_int_TestPlace.Col16; []int <= []bool")
	}
}

/*

	uint64 test
*/

type array_uint64_TestPlace struct {
	Col1  int      `db:"col1" type:"int"`
	Col2  []uint64 `db:"col2" type:"[]bigint"`
	Col3  []uint64 `db:"col3" type:"[]smallint"`
	Col4  []uint64 `db:"col4" type:"[]int"`
	Col7  []uint64 `db:"col7" type:"[]text"`
	Col8  []uint64 `db:"col8" type:"[]varchar"`
	Col9  []uint64 `db:"col9" type:"[]char"`
	Col11 []uint64 `db:"col11" type:"[]real"`
	Col12 []uint64 `db:"col12" type:"[]double"`
	Col13 []uint64 `db:"col13" type:"[]numeric"`
	Col14 []uint64 `db:"col14" type:"[]decimal"`
	Col15 []uint64 `db:"col15" type:"[]money"`
	Col16 []uint64 `db:"col16" type:"[]bool"`
}

var array_uint64_mTestType *AllRows = &AllRows{
	SType: reflect.TypeOf(array_uint64_TestPlace{}),
}

func _03_array_uint64(t *testing.T, s *Searcher) {
	array_main_f_test_table(s)
	p := []array_uint64_TestPlace{}
	s.Get(array_uint64_mTestType, &p, "SELECT * FROM public.test ORDER BY 1")

	if p[0].Col2[0] != 9223372036854775807 {
		t.Fatal("Error array_int_TestPlace.Col2; []int <= []bigint")
	}
	if p[0].Col3[0] != 18 {
		t.Fatal("Error array_int_TestPlace.Col3; []int <= []smallint")
	}
	if p[0].Col4[0] != 884 {
		t.Fatal("Error array_int_TestPlace.Col4; []int <= []int")
	}
	if p[0].Col7[0] != 10 {
		t.Fatal("Error array_int_TestPlace.Col7; []int <= []text")
	}
	if p[0].Col8[1] != 0 {
		t.Fatal("Error array_int_TestPlace.Col8; []int <= []varchar")
	}
	if p[0].Col16[0] != 1 {
		t.Fatal("Error array_int_TestPlace.Col16; []int <= []bool")
	}
}

/*

	uint test
*/

type array_uint_TestPlace struct {
	Col1  int    `db:"col1" type:"int"`
	Col2  []uint `db:"col2" type:"[]bigint"`
	Col3  []uint `db:"col3" type:"[]smallint"`
	Col4  []uint `db:"col4" type:"[]int"`
	Col7  []uint `db:"col7" type:"[]text"`
	Col8  []uint `db:"col8" type:"[]varchar"`
	Col9  []uint `db:"col9" type:"[]char"`
	Col11 []uint `db:"col11" type:"[]real"`
	Col12 []uint `db:"col12" type:"[]double"`
	Col13 []uint `db:"col13" type:"[]numeric"`
	Col14 []uint `db:"col14" type:"[]decimal"`
	Col15 []uint `db:"col15" type:"[]money"`
	Col16 []uint `db:"col16" type:"[]bool"`
}

var array_uint_mTestType *AllRows = &AllRows{
	SType: reflect.TypeOf(array_uint_TestPlace{}),
}

func _04_array_uint(t *testing.T, s *Searcher) {
	array_main_f_test_table(s)
	p := []array_uint_TestPlace{}
	s.Get(array_uint_mTestType, &p, "SELECT * FROM public.test ORDER BY 1")

	if p[0].Col2[0] != 9223372036854775807 {
		t.Fatal("Error array_int_TestPlace.Col2; []int <= []bigint")
	}
	if p[0].Col3[0] != 18 {
		t.Fatal("Error array_int_TestPlace.Col3; []int <= []smallint")
	}
	if p[0].Col4[0] != 884 {
		t.Fatal("Error array_int_TestPlace.Col4; []int <= []int")
	}
	if p[0].Col7[0] != 10 {
		t.Fatal("Error array_int_TestPlace.Col7; []int <= []text")
	}
	if p[0].Col8[1] != 0 {
		t.Fatal("Error array_int_TestPlace.Col8; []int <= []varchar")
	}
	if p[0].Col16[0] != 1 {
		t.Fatal("Error array_int_TestPlace.Col16; []int <= []bool")
	}
}

/*
	boolean test
*/
type array_bool_TestPlace struct {
	Col1  int       `db:"col1" type:"int"`
	Col2  []string  `db:"col2" type:"[]boolean"`
	Col3  []int64   `db:"col3" type:"[]boolean"`
	Col4  []float32 `db:"col4" type:"[]boolean"`
	Col5  []float64 `db:"col5" type:"[]bool"`
	Col6  []bool    `db:"col6" type:"[]bool"`
	Col7  []int     `db:"col7" type:"[]bool"`
	Col8  []uint8   `db:"col8" type:"[]bool"`
	Col9  []uint64  `db:"col9" type:"[]bool"`
	Col10 []uint    `db:"col10" type:"[]bool"`
}

var array_bool_mTestType *AllRows = &AllRows{
	SType: reflect.TypeOf(array_bool_TestPlace{}),
}

func _11_array_bool(t *testing.T, s *Searcher) {

	sql_create := " CREATE TABLE public.test " +
		"(col1 int, col2 boolean[], col3 boolean[], col4 boolean[], " +
		"col5 boolean[], col6 boolean[], col7 boolean[], col8 boolean[], col9 boolean[], col10 boolean[]) "
	sql_cols := "INSERT INTO test(col1, col2, col3, col4, col5, col6, col7, col8, col9, col10 ) "
	sql_vals := []string{
		"VALUES (1, '{TRUE,FALSE}'::bool[], '{TRUE,FALSE}'::bool[], '{TRUE,FALSE}'::bool[], '{TRUE,FALSE}'::bool[], '{TRUE,FALSE}'::bool[], '{TRUE,FALSE}'::bool[], '{TRUE,FALSE}'::bool[],'{TRUE,FALSE}'::bool[],'{TRUE,FALSE}'::bool[] )",
		"VALUES (2, '{FALSE,TRUE}'::bool[], '{FALSE,TRUE}'::bool[], '{FALSE,TRUE}'::bool[], '{FALSE,TRUE}'::bool[], '{FALSE,TRUE}'::bool[], '{FALSE,TRUE}'::bool[], '{FALSE,TRUE}'::bool[], '{FALSE,TRUE}'::bool[], '{FALSE,TRUE}'::bool[])",
		"VALUES (3, null,  null,  null,  null,  null,  null,  null,  null, null)", // check null - nil
	}

	make_t_table(s, sql_create, sql_cols, sql_vals)
	p := []array_bool_TestPlace{}
	s.Get(array_bool_mTestType, &p, "SELECT * FROM public.test ORDER BY 1")

	if p[0].Col2[0] != "1" {
		t.Fatal("Error array_int_TestPlace.Col2; []string <= []boolean")
	}
	if p[0].Col3[0] != 1 {
		t.Fatal("Error array_int_TestPlace.Col3; []int64 <= []boolean")
	}
	if p[0].Col4[0] != 0 {
		t.Fatal("Error array_int_TestPlace.Col4; []float32 <= []boolean")
	}
	if p[0].Col5[0] != 0 {
		t.Fatal("Error array_int_TestPlace.Col5; []float64 <= []boolean")
	}
	if p[0].Col6[0] != true {
		t.Fatal("Error array_int_TestPlace.Col6; []bool <= []boolean")
	}
	if p[0].Col7[0] != 1 {
		t.Fatal("Error array_int_TestPlace.Col7; []int <= []boolean")
	}
	if p[0].Col8[1] != 0 {
		t.Fatal("Error array_int_TestPlace.Col8; []uint8 <= []boolean")
	}
	if p[0].Col9[0] != 1 {
		t.Fatal("Error array_int_TestPlace.Col16; []uint64 <= []boolean")
	}
}

/*
	boolean test
*/
type array_bool2_TestPlace struct {
	Col1  int    `db:"col1" type:"int"`
	Col2  []bool `db:"col2" type:"[]bigint"`
	Col3  []bool `db:"col3" type:"[]smallint"`
	Col4  []bool `db:"col4" type:"[]int"`
	Col5  []bool `db:"col5" type:"[]text"`
	Col6  []bool `db:"col6" type:"[]varchar"`
	Col7  []bool `db:"col7" type:"[]char"`
	Col8  []bool `db:"col8" type:"[]real"`
	Col9  []bool `db:"col9" type:"[]double"`
	Col10 []bool `db:"col10" type:"[]numeric"`
	Col11 []bool `db:"col11" type:"[]decimal"`
	Col12 []bool `db:"col12" type:"[]money"`
	Col13 []bool `db:"col13" type:"[]bool"`
}

func check_bool_result(t *testing.T, col, p []bool, name string) {
	if len(col) != len(p) {
		t.Fatalf("\n%s: %#v\np: %#v\n, Error _12_array_bool Bad len for %s [%d <=> %d] ",
			name, col, p, name, len(col), len(p))
	}
	for i := range col {
		if p[i] != col[i] {
			t.Fatalf("\n%s: %#v\np: %#v\nError _12_array_bool Bad value for %s[%d] [%s <=> %s]",
				name, col, p, name, i, p[i], col[i])
		}
	}
}

func _12_array_bool(t *testing.T, s *Searcher) {

	sql_create := " CREATE TABLE public.test " +
		"(col1 int, col2 bigint[], col3 smallint[], col4 integer[], " +
		" col5 text[], col6 varchar(50)[], col7 char(10)[], " +
		" col8 real[], col9 double precision[], col10 numeric[], col11 decimal[], " +
		" col12 money[], col13 boolean[]  " +
		") "

	sql_cols := "INSERT INTO test(col1, col2, col3, col4, col5, col6, col7, col8, col9, col10, col11, col12, col13 ) "
	sql_vals := []string{" VALUES (1, " +
		// cols 2,3,4
		"'{9223372036854775807,3,2,0}'::bigint[],  '{18,28,0,33}'::smallint[],  '{884,-121,0}'::int[], " +
		// cols 5,6
		" '{\"10\",\"\",\"Малененький текст\"}'::text[], '{\"\", \" varchar next\", \"654\"}', " +
		// cols 7,8
		"'{\" char[] \",\"\",\" char2[] \",\"123.56234\",\"\"}', '{0,-12.13,13.13213409,0,-12.130909,0}'," +
		// cols 9, 10
		"'{14.105,0,-15.015,-15.1500}', '{0,-16.17,0.0,16.1600,0,-16.173244,0}', " +
		// cols 11,12
		"'{18.00,18.18,1800.180999234,0.0}'::decimal[], '{21.21,22.22,0.0,0,23.23}'::money[]," +
		// cols 13
		" '{TRUE,FALSE,TRUE,FALSE}'::boolean[] ) ",
		" VALUES (2, null, null, null, null, null, null, null, null, null, null, null, null ) ", // check null - nil
	}

	make_t_table(s, sql_create, sql_cols, sql_vals)

	var array_bool_TestPlace *AllRows = &AllRows{
		SType: reflect.TypeOf(array_bool2_TestPlace{}),
	}

	col2 := []bool{true, true, true, false}
	col3 := []bool{true, true, false, true}
	col4 := []bool{true, true, false}
	col5 := []bool{true, false, true}
	col6 := []bool{false, true, true}
	col7 := []bool{true, false, true, true, false}
	col8 := []bool{false, true, true, false, true, false}
	col9 := []bool{true, false, true, true}
	col10 := []bool{false, true, false, true, false, true, false}
	col11 := []bool{true, true, true, false}
	col12 := []bool{true, true, false, false, true}
	col13 := []bool{true, false, true, false}

	p := []array_bool2_TestPlace{}
	s.Get(array_bool_TestPlace, &p, "SELECT * FROM public.test ORDER BY 1")

	check_bool_result(t, col2, p[0].Col2, "col2")
	check_bool_result(t, col3, p[0].Col3, "col3")
	check_bool_result(t, col4, p[0].Col4, "col4")
	check_bool_result(t, col5, p[0].Col5, "col5")
	check_bool_result(t, col6, p[0].Col6, "col6")
	check_bool_result(t, col7, p[0].Col7, "col7")
	check_bool_result(t, col8, p[0].Col8, "col8")
	check_bool_result(t, col9, p[0].Col9, "col9")
	check_bool_result(t, col10, p[0].Col10, "col10")
	check_bool_result(t, col11, p[0].Col11, "col11")
	check_bool_result(t, col12, p[0].Col12, "col12")
	check_bool_result(t, col13, p[0].Col13, "col13")

}

/*
	string test
*/
type array_string_TestPlace struct {
	Col1  int      `db:"col1" type:"int"`
	Col2  []string `db:"col2" type:"[]bigint"`
	Col3  []string `db:"col3" type:"[]smallint"`
	Col4  []string `db:"col4" type:"[]integer"`
	Col7  []string `db:"col7" type:"[]text"`
	Col8  []string `db:"col8" type:"[]varchar"`
	Col9  []string `db:"col9" type:"[]char"`
	Col11 []string `db:"col11" type:"[]real"`
	Col12 []string `db:"col12" type:"[]double"`
	Col13 []string `db:"col13" type:"[]numeric"`
	Col14 []string `db:"col14" type:"[]decimal"`
	Col15 []string `db:"col15" type:"[]money"`
	Col16 []string `db:"col16" type:"[]bool"`
}

var array_string_mTestType *AllRows = &AllRows{
	SType: reflect.TypeOf(array_string_TestPlace{}),
}

func _21_array_string(t *testing.T, s *Searcher) {
	array_main_f_test_table(s)
	p := []array_string_TestPlace{}
	s.Get(array_string_mTestType, &p, "SELECT * FROM public.test ORDER BY 1")

	if p[0].Col2[0] != "9223372036854775807" {
		t.Fatal("Error array_int_TestPlace.Col2; []string <= []bigint")
	}
	if p[0].Col3[0] != "18" {
		t.Fatal("Error array_int_TestPlace.Col3; []string <= []smallint")
	}
	if p[0].Col4[0] != "884" {
		t.Fatal("Error array_int_TestPlace.Col4; []string <= []int")
	}
	if p[0].Col7[0] != "10" {
		t.Fatal("Error array_int_TestPlace.Col7; []string <= []text")
	}
	if p[0].Col8[1] != " varchar next" {
		t.Fatal("Error array_int_TestPlace.Col8; []string <= []varchar")
	}
	if p[0].Col16[0] != "1" {
		t.Fatal("Error array_int_TestPlace.Col16; []string <= []bool")
	}
}

/*
	float64 test
*/
type array_float32_TestPlace struct {
	Col1  int       `db:"col1" type:"int"`
	Col2  []float32 `db:"col2" type:"[]bigint"`
	Col3  []float32 `db:"col3" type:"[]smallint"`
	Col4  []float32 `db:"col4" type:"[]integer"`
	Col7  []float32 `db:"col7" type:"[]text"`
	Col8  []float32 `db:"col8" type:"[]varchar"`
	Col9  []float32 `db:"col9" type:"[]char"`
	Col11 []float32 `db:"col11" type:"[]real"`
	Col12 []float32 `db:"col12" type:"[]double"`
	Col13 []float32 `db:"col13" type:"[]numeric"`
	Col14 []float32 `db:"col14" type:"[]decimal"`
	Col15 []float32 `db:"col15" type:"[]money"`
	Col16 []float32 `db:"col16" type:"[]bool"`
}

var array_float32_mTestType *AllRows = &AllRows{
	SType: reflect.TypeOf(array_float32_TestPlace{}),
}

func _31_array_float32(t *testing.T, s *Searcher) {
	array_main_f_test_table(s)
	p := []array_float32_TestPlace{}
	s.Get(array_float32_mTestType, &p, "SELECT * FROM public.test ORDER BY 1")

	if p[0].Col2[0] != 9223372036854775807 {
		t.Fatal("Error array_int_TestPlace.Col2; []float32 <= []bigint")
	}
	if p[0].Col3[0] != 18 {
		t.Fatal("Error array_int_TestPlace.Col3; []float32 <= []smallint")
	}
	if p[0].Col4[0] != 884 {
		t.Fatal("Error array_int_TestPlace.Col4; []float32 <= []int")
	}
	if p[0].Col7[0] != 10 {
		t.Fatal("Error array_int_TestPlace.Col7; []float32 <= []text")
	}
	if p[0].Col8[0] != 0 {
		t.Fatal("Error array_int_TestPlace.Col8; []float32 <= []varchar")
	}
	if p[0].Col16[0] != 0 {
		t.Fatal("Error array_int_TestPlace.Col16; []float32 <= []bool")
	}
}

/*
	float64 test
*/
type array_float64_TestPlace struct {
	Col1  int       `db:"col1" type:"int"`
	Col2  []float64 `db:"col2" type:"[]bigint"`
	Col3  []float64 `db:"col3" type:"[]smallint"`
	Col4  []float64 `db:"col4" type:"[]integer"`
	Col7  []float64 `db:"col7" type:"[]text"`
	Col8  []float64 `db:"col8" type:"[]varchar"`
	Col9  []float64 `db:"col9" type:"[]char"`
	Col11 []float64 `db:"col11" type:"[]real"`
	Col12 []float64 `db:"col12" type:"[]double"`
	Col13 []float64 `db:"col13" type:"[]numeric"`
	Col14 []float64 `db:"col14" type:"[]decimal"`
	Col15 []float64 `db:"col15" type:"[]money"`
	Col16 []float64 `db:"col16" type:"[]bool"`
}

var array_float64_mTestType *AllRows = &AllRows{
	SType: reflect.TypeOf(array_float64_TestPlace{}),
}

func _32_array_float64(t *testing.T, s *Searcher) {
	array_main_f_test_table(s)
	p := []array_float64_TestPlace{}
	s.Get(array_float64_mTestType, &p, "SELECT * FROM public.test ORDER BY 1")

	if p[0].Col2[0] != 9223372036854775807 {
		t.Fatal("Error array_int_TestPlace.Col2; []float64 <= []bigint")
	}
	if p[0].Col3[0] != 18 {
		t.Fatal("Error array_int_TestPlace.Col3; []float64 <= []smallint")
	}
	if p[0].Col4[0] != 884 {
		t.Fatal("Error array_int_TestPlace.Col4; []float64 <= []int")
	}
	if p[0].Col7[0] != 10 {
		t.Fatal("Error array_int_TestPlace.Col7; []float64 <= []text")
	}
	if p[0].Col8[0] != 0 {
		t.Fatal("Error array_int_TestPlace.Col8; []float64 <= []varchar")
	}
	if p[0].Col16[0] != 0 {
		t.Fatal("Error array_int_TestPlace.Col16; []float64 <= []bool")
	}
}
