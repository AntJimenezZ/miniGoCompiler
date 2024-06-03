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
	%4 = icmp sle i32 %3, 20
	br i1 %4, label %5, label %7

5:
	%6 = call i32 (i8*, ...) @printf([4 x i8]* @0, i32 1)
	br label %9

7:
	%8 = call i32 (i8*, ...) @printf([4 x i8]* @0, i32 2)
	br label %9

9:
	%10 = load i32, i32* %1
	%11 = icmp sge i32 %10, 50
	br i1 %11, label %12, label %14

12:
	%13 = call i32 (i8*, ...) @printf([4 x i8]* @0, i32 23)
	br label %27

14:
	%15 = load i32, i32* %1
	%16 = icmp eq i32 %15, 10
	br i1 %16, label %17, label %27

17:
	%18 = mul i32 5, 2
	%19 = add i32 %18, 2
	%20 = sdiv i32 8, 2
	%21 = mul i32 %20, 4
	%22 = sub i32 %19, %21
	%23 = alloca i32
	store i32 %22, i32* %23
	%24 = load i32, i32* %23
	%25 = load i32, i32* %23
	%26 = call i32 (i8*, ...) @printf([4 x i8]* @0, i32 %25)
	br label %27

27:
	ret i32 0
}
