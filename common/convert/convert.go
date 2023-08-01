package convert

func GroupCvtChinese(group string) string {
	switch group {
	case "Backend":
		group = "后端组"
	case "Product":
		group = "产品组"
	case "Design":
		group = "设计组"
	case "Android":
		group = "安卓组"
	case "Frontend":
		group = "前端组"
	}
	return group
}
