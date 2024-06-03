@0 = global [4 x i8] c"%d\0A\00"
@1 = global [4 x i8] c"%f\0A\00"
@2 = global [4 x i8] c"%s\0A\00"
@3 = global [4 x i8] c"%c\0A\00"
@global = global i32 10
@a = global i32 34
@b = global [6 x i8] c"\22Hola\22"
@c = global float 30.5

declare i32 @printf(i8* %0, ...)

define i32 @suma(i32 %0) {
1:
	%2 = alloca i32
	store i32 %0, i32* %2
	store i32 25, i32* %2
	%3 = load i32, i32* %2
	%4 = add i32 %3, 10
	%5 = alloca i32
	store i32 %4, i32* %5
	%6 = load i32, i32* %2
	%7 = add i32 %6, 10
	%8 = load i32, i32* %5
	ret i32 %8
}

define i32 @main() {
0:
	%1 = alloca [10 x i32]
	store [10 x i32] zeroinitializer, [10 x i32]* %1
	%2 = call i32 (i8*, ...) @printf([4 x i8]* @0, i32 10)
	%3 = alloca i32
	store i32 10, i32* %3
	%4 = load i32, i32* %3
	%5 = load i32, i32* %3
	%6 = icmp sle i32 %5, 9
	br i1 %6, label %7, label %9

7:
	%8 = call i32 (i8*, ...) @printf([4 x i8]* @0, i32 1)
	br label %11

9:
	%10 = call i32 (i8*, ...) @printf([4 x i8]* @0, i32 2)
	br label %11

11:
	br label %12

12:
	%13 = call i32 (i8*, ...) @printf([4 x i8]* @0, i32 1111)
	br label %12

14:
	%15 = load i32, i32* %3
	%16 = icmp eq i32 %15, 11
	br i1 %16, label %17, label %19

17:
	%18 = call i32 (i8*, ...) @printf([4 x i8]* @0, i32 23)
	br label %23

19:
	%20 = call i32 (i8*, ...) @printf([4 x i8]* @0, i32 0)
	%21 = load i32, i32* %3
	%22 = icmp sge i32 %21, 11
	br i1 %22, label %19, label %23

23:
	ret i32 0
}
