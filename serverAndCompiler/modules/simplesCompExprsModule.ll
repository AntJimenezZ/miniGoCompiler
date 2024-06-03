@0 = global [4 x i8] c"%d\0A\00"
@1 = global [4 x i8] c"%f\0A\00"
@2 = global [4 x i8] c"%s\0A\00"
@3 = global [4 x i8] c"%c\0A\00"

declare i32 @printf(i8* %0, ...)

define i32 @main() {
0:
	%1 = add i32 1, 2
	%2 = alloca i32
	store i32 %1, i32* %2
	%3 = load i32, i32* %2
	%4 = sub i32 3, 2
	%5 = alloca i32
	store i32 %4, i32* %5
	%6 = load i32, i32* %5
	%7 = mul i32 3, 3
	%8 = alloca i32
	store i32 %7, i32* %8
	%9 = load i32, i32* %8
	%10 = sdiv i32 2, 2
	%11 = alloca i32
	store i32 %10, i32* %11
	%12 = load i32, i32* %11
	%13 = srem i32 2, 10
	%14 = alloca i32
	store i32 %13, i32* %14
	%15 = load i32, i32* %14
	%16 = icmp sgt i32 10, 12
	%17 = alloca i1
	store i1 %16, i1* %17
	%18 = load i1, i1* %17
	%19 = mul i32 10, 2
	%20 = mul i32 %19, 5
	%21 = add i32 5, 2
	%22 = sdiv i32 2, %21
	%23 = sub i32 %20, %22
	%24 = alloca i32
	store i32 %23, i32* %24
	%25 = load i32, i32* %24
	%26 = load i32, i32* %2
	%27 = load i32, i32* %5
	%28 = icmp ne i32 %26, %27
	%29 = alloca i1
	store i1 %28, i1* %29
	%30 = load i1, i1* %29
	%31 = load i32, i32* %2
	%32 = icmp eq i32 %31, 3
	%33 = alloca i1
	store i1 %32, i1* %33
	%34 = load i1, i1* %33
	%35 = load i32, i32* %8
	%36 = icmp sle i32 %35, 8
	%37 = alloca i1
	store i1 %36, i1* %37
	%38 = load i1, i1* %37
	%39 = load i32, i32* %2
	%40 = call i32 (i8*, ...) @printf([4 x i8]* @0, i32 %39)
	%41 = load i32, i32* %5
	%42 = call i32 (i8*, ...) @printf([4 x i8]* @0, i32 %41)
	%43 = load i32, i32* %8
	%44 = call i32 (i8*, ...) @printf([4 x i8]* @0, i32 %43)
	%45 = load i32, i32* %11
	%46 = call i32 (i8*, ...) @printf([4 x i8]* @0, i32 %45)
	%47 = load i32, i32* %14
	%48 = call i32 (i8*, ...) @printf([4 x i8]* @0, i32 %47)
	%49 = load i1, i1* %17
	%50 = call i32 (i8*, ...) @printf([4 x i8]* @0, i1 %49)
	%51 = load i1, i1* %29
	%52 = call i32 (i8*, ...) @printf([4 x i8]* @0, i1 %51)
	%53 = load i1, i1* %33
	%54 = call i32 (i8*, ...) @printf([4 x i8]* @0, i1 %53)
	%55 = load i32, i32* %24
	%56 = call i32 (i8*, ...) @printf([4 x i8]* @0, i32 %55)
	%57 = load i1, i1* %37
	%58 = call i32 (i8*, ...) @printf([4 x i8]* @0, i1 %57)
	ret i32 0
}
