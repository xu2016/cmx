package xjhcm

const (
	nstr    = "0123456789"
	sdstr   = "klmnopqrwxyzabstuvghijcdef"
	sustr   = "KLMNOPQRWXYZABSTUVGHIJCDEF"
	sastr   = "klmnopqrwxyzabstuvghijcdefKLMNOPQRWXYZABSTUVGHIJCDEF"
	nsdstr  = "kl3mn2opqr1wx9yz8abst5uv4gh7ij6cd0ef"
	nsustr  = "KL3MN2OPQR1WX9YZ8ABST5UV4GH7IJ6CD0EF"
	nsastr  = "kl0mnop9qrwxy7zabs5tuvgh8ijcdefK6LMNOP1QRWX3YZABS4TUVGHI2JCDEF"
	codestr = "0123456789klmnopqrwxyzabstuvghijcdefKLMNOPQRWXYZABSTUVGHIJCDEF"
	keystr  = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

//设置生成字符串的格式
const (
	NSTR   = iota //数字字符串
	SDSTR         //小写字母字符串
	SUSTR         //大写字母字符串
	SASTR         //大写和小写字母字符串
	NSDSTR        //数字和小写字母字符串
	NSUSTR        //数字和大写字母字符串
	NSASTR        //数字大写和小写字母字符串
	KEYSTR        //数字大写和小写字母字符串(有序)
)

var num = map[string]int{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9}
