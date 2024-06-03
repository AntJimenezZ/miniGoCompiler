@0 = global [4 x i8] c"%d\0A\00"
@1 = global [4 x i8] c"%f\0A\00"
@2 = global [4 x i8] c"%s\0A\00"
@3 = global [4 x i8] c"%c\0A\00"
@a = global i32 zeroinitializer
@b = global i32 10
@c = global i32 11
@d = global float zeroinitializer
@e = global float 1.0
@f = global float 10.0
@g = global i8* zeroinitializer
@h = global [6 x i8] c"\22hola\22"
@i = global [7 x i8] c"\22kevin\22"
@j = global i1 zeroinitializer
@k = global i32 zeroinitializer
@l = global i32 12
@m = global i32 15
@n = global float zeroinitializer
@o = global float 15.0
@p = global float 2.0
@q = global i8* zeroinitializer
@r = global [5 x i8] c"\22ver\22"
@s = global [6 x i8] c"\22aaaa\22"
@t = global float 1.0
@u = global i32 2
@v = global [7 x i8] c"\22carro\22"

declare i32 @printf(i8* %0, ...)

define i32 @main() {
0:
	%1 = alloca [8 x i8]
	store [8 x i8] c"\22Fabian\22", [8 x i8]* %1
	%2 = alloca [6 x i8]
	store [6 x i8] c"\22hola\22", [6 x i8]* %2
	%3 = load [6 x i8], [6 x i8]* %2
	%4 = alloca i8*
	store i8* zeroinitializer, i8** %4
	ret i32 0
}
