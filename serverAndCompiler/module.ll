@0 = global [4 x i8] c"%d\0A\00"
@1 = global [4 x i8] c"%f\0A\00"
@2 = global [4 x i8] c"%s\0A\00"
@3 = global [4 x i8] c"%c\0A\00"

declare i32 @printf(i8* %0, ...)

define i32 @caca() {
0:
	%1 = mul i32 10, 2
	%2 = alloca i32
	store i32 %1, i32* %2
	%3 = mul i32 10, 2
	br label %4

4:
	%5 = alloca i32
	store i32 zeroinitializer, i32* %5
	%6 = alloca i32
	store i32 zeroinitializer, i32* %6
	%7 = alloca i32
	store i32 zeroinitializer, i32* %7
	%8 = alloca i32
	store i32 zeroinitializer, i32* %8
	ret i32 0
}

define i32 @main() {
0:
	%1 = alloca i32
	store i32 10, i32* %1
	%2 = load i32, i32* %1
	%3 = icmp eq i32 %2, 10
	br i1 %3, label %4, label %6

4:
	%5 = alloca i32
	store i32 zeroinitializer, i32* %5
	br label %6

6:
	ret i32 0
}
