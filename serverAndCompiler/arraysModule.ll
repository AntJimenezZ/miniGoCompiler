@0 = global [4 x i8] c"%d\0A\00"
@1 = global [4 x i8] c"%f\0A\00"
@2 = global [4 x i8] c"%s\0A\00"
@3 = global [4 x i8] c"%c\0A\00"
@array1 = global [5 x i32] zeroinitializer
@array2 = global [12 x i8*] zeroinitializer

declare i32 @printf(i8* %0, ...)

define i32 @main() {
0:
	%1 = alloca [1 x float]
	store [1 x float] zeroinitializer, [1 x float]* %1
	ret i32 0
}
