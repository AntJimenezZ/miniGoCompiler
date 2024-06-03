@0 = global [4 x i8] c"%d\0A\00"
@1 = global [4 x i8] c"%f\0A\00"
@2 = global [4 x i8] c"%s\0A\00"
@3 = global [4 x i8] c"%c\0A\00"

declare i32 @printf(i8* %0, ...)

define i32 @hola(i32 %0, float %1) {
2:
	%3 = alloca i32
	store i32 %0, i32* %3
	%4 = alloca float
	store float %1, float* %4
	store i32 10, i32* %3
	store float 12.5, float* %4
	ret i32 0
}

define float @compiladores() {
0:
	ret float 10.0
}

define void @amigo(i32 %0, i32 %1) {
2:
	%3 = alloca i32
	store i32 %0, i32* %3
	%4 = alloca i32
	store i32 %1, i32* %4
	store i32 10, i32* %3
	store i32 11, i32* %4
	ret void
}

define i32 @main() {
0:
	%1 = alloca [1 x float]
	store [1 x float] zeroinitializer, [1 x float]* %1
	ret i32 0
}
