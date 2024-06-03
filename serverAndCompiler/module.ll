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
	br label %1

1:
	%2 = alloca i32
	store i32 0, i32* %2
	br label %3

3:
	%4 = load i32, i32* %2
	%5 = add i32 %4, 1
	store i32 %5, i32* %2
	%6 = load i32, i32* %2
	%7 = call i32 (i8*, ...) @printf([4 x i8]* @0, i32 %6)
	br label %1

8:
	ret i32 0
}
