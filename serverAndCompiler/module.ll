@0 = global [4 x i8] c"%d\0A\00"
@1 = global [4 x i8] c"%f\0A\00"
@2 = global [4 x i8] c"%s\0A\00"
@3 = global [4 x i8] c"%c\0A\00"
@arr = global [5 x i32] zeroinitializer

declare i32 @printf(i8* %0, ...)

define i32 @main() {
0:
	%1 = call i32 (i8*, ...) @printf([4 x i8]* @0, i32 2)
	%2 = mul i32 10, 2
	%3 = alloca i32
	store i32 %2, i32* %3
	%4 = mul i32 10, 2
	%5 = load i32, i32* %3
	%6 = icmp eq i32 %5, 10
	br i1 %6, label %7, label %9

7:
	%8 = alloca i32
	store i32 zeroinitializer, i32* %8
	br label %9

9:
	ret i32 0
}
