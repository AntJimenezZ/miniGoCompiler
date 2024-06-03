@0 = global [4 x i8] c"%d\0A\00"
@1 = global [4 x i8] c"%f\0A\00"
@2 = global [4 x i8] c"%s\0A\00"
@3 = global [4 x i8] c"%c\0A\00"

declare i32 @printf(i8* %0, ...)

define i32 @main() {
0:
	%1 = alloca i32
	store i32 10, i32* %1
	%2 = load i32, i32* %1
	%3 = load i32, i32* %1
	%4 = icmp slt i32 %3, 12
	br i1 %4, label %5, label %8

5:
	%6 = load i32, i32* %1
	%7 = call i32 (i8*, ...) @printf([4 x i8]* @0, i32 %6)
	br label %8

8:
	ret i32 0
}
