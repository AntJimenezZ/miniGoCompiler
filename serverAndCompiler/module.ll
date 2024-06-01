@0 = global [4 x i8] c"%d\0A\00"
@1 = global [4 x i8] c"%f\0A\00"
@2 = global [4 x i8] c"%s\0A\00"
@3 = global [4 x i8] c"%c\0A\00"
@x = global i32 10

declare i32 @printf(i8* %0, ...)

define i32 @main(i32 %0, i32 %1) {
2:
	%3 = alloca i32
	store i32 %0, i32* %3
	%4 = alloca i32
	store i32 %1, i32* %4
	store i32 10, i32* %3
	%5 = alloca i32
	store i32 10, i32* %5
	%6 = load i32, i32* %3
	%7 = icmp eq i32 %6, 10
	br i1 %7, label %8, label %10

8:
	%9 = alloca i32
	store i32 20, i32* %9
	br label %10

10:
	ret i32 0
}
