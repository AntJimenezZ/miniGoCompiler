@0 = global [4 x i8] c"%d\0A\00"
@1 = global [4 x i8] c"%f\0A\00"
@2 = global [4 x i8] c"%s\0A\00"
@3 = global [4 x i8] c"%c\0A\00"
@global = global i32 20
@palabra = global i8* zeroinitializer
@puntoFlotante = global float 0x4024666660000000
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

define void @imprimir() {
0:
	%1 = call i32 (i8*, ...) @printf([4 x i8]* @0, i32 55)
	ret void
}

define i32 @main() {
0:
	%1 = icmp ne i32 1, 2
	%2 = alloca i1
	store i1 %1, i1* %2
	%3 = load i1, i1* %2
	%4 = alloca [10 x i32]
	store [10 x i32] zeroinitializer, [10 x i32]* %4
	%5 = alloca i32
	store i32 10, i32* %5
	%6 = load i32, i32* %5
	%7 = call i32 (i8*, ...) @printf([4 x i8]* @0, i32 %6)
	%8 = alloca i32
	store i32 10, i32* %8
	%9 = load i32, i32* %8
	%10 = load i32, i32* %8
	%11 = icmp sle i32 %10, 20
	br i1 %11, label %12, label %14

12:
	%13 = call i32 (i8*, ...) @printf([4 x i8]* @0, i32 1)
	br label %16

14:
	%15 = call i32 (i8*, ...) @printf([4 x i8]* @0, i32 2)
	br label %16

16:
	%17 = load i32, i32* %8
	%18 = icmp sge i32 %17, 50
	br i1 %18, label %19, label %21

19:
	%20 = call i32 (i8*, ...) @printf([4 x i8]* @0, i32 23)
	br label %34

21:
	%22 = load i32, i32* %8
	%23 = icmp eq i32 %22, 10
	br i1 %23, label %24, label %34

24:
	%25 = mul i32 5, 2
	%26 = add i32 %25, 2
	%27 = sdiv i32 8, 2
	%28 = mul i32 %27, 4
	%29 = sub i32 %26, %28
	%30 = alloca i32
	store i32 %29, i32* %30
	%31 = load i32, i32* %30
	%32 = load i32, i32* %30
	%33 = call i32 (i8*, ...) @printf([4 x i8]* @0, i32 %32)
	br label %34

34:
	ret i32 0
}
