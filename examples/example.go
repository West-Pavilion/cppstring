package main

import (
	"fmt"
	"github.com/West-Pavilion/cppstring"
)

func main() {
	fmt.Println("=== cppstring 库使用示例 ===")
	fmt.Println("=== cppstring Library Usage Example ===")
	fmt.Println()

	// 测试基本英文字符串
	fmt.Println("1. 基本英文字符串测试 (Basic English String Test):")
	english := "Hello World"
	fmt.Printf("   原字符串 (Original): %s\n", english)
	fmt.Printf("   反转结果 (Reversed): %s\n", cppstring.Reverse(english))
	fmt.Println()

	// 测试中文字符串
	fmt.Println("2. 中文字符串测试 (Chinese String Test):")
	chinese := "你好世界"
	fmt.Printf("   原字符串 (Original): %s\n", chinese)
	fmt.Printf("   反转结果 (Reversed): %s\n", cppstring.Reverse(chinese))
	fmt.Println()

	// 测试混合字符串
	fmt.Println("3. 混合字符串测试 (Mixed String Test):")
	mixed := "Hello, 世界! 123"
	fmt.Printf("   原字符串 (Original): %s\n", mixed)
	fmt.Printf("   反转结果 (Reversed): %s\n", cppstring.Reverse(mixed))
	fmt.Println()

	// 测试表情符号
	fmt.Println("4. 表情符号测试 (Emoji Test):")
	emoji := "Hello 👋 World 🌍"
	fmt.Printf("   原字符串 (Original): %s\n", emoji)
	fmt.Printf("   反转结果 (Reversed): %s\n", cppstring.Reverse(emoji))
	fmt.Println()

	// 测试空字符串
	fmt.Println("5. 空字符串测试 (Empty String Test):")
	empty := ""
	fmt.Printf("   原字符串 (Original): '%s'\n", empty)
	fmt.Printf("   反转结果 (Reversed): '%s'\n", cppstring.Reverse(empty))
	fmt.Println()

	fmt.Println("=== 测试完成 (Tests Completed) ===")
}