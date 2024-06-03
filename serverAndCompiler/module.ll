@0 = global [4 x i8] c"%d\0A\00"
@1 = global [4 x i8] c"%f\0A\00"
@2 = global [4 x i8] c"%s\0A\00"
@3 = global [4 x i8] c"%c\0A\00"

declare i32 @printf(i8* %0, ...)

define i32 @caca(i32 %0) {
1:
	%2 = alloca i32
	store i32 %0, i32* %2
	store i32 10, i32* %2
	%3 = call i32 (i8*, ...) @printf([4 x i8]* @0, i32 10)
	ret i32 0
}

define i32 @main() {
0:
	%1 = alloca [5 x i32]
	store [5 x i32] zeroinitializer, [5 x i32]* %1
	%2 = alloca i32
	store i32 5, i32* %2
	%3 = call i32 (i8*, ...) @printf([4 x i8]* @0, i32 5)
	%4 = icmp slt i32 3, 5
	%5 = alloca i1
	store i1 %4, i1* %5
	%6 = load i1, i1* %5
	%7 = load i1, i1* %5
	%8 = alloca i32
	store i32 10, i32* %8
	%9 = load i32, i32* %8
	%10 = icmp eq i32 %9, 10
	br i1 %10, label %11, label %14

11:
	%12 = alloca float
	store float 10.0, float* %12
	%13 = call i32 (i8*, ...) @printf([4 x i8]* @1, float 10.0)
	br label %14

14:
	ret i32 0
}
