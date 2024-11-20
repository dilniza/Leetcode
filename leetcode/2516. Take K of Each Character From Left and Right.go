package leetcode

func TakeCharacters(s string, k int) int {
    n := len(s) // n - satr uzunligini aniqlash
    if k == 0 { 
        return 0 // Agar `k == 0` bo'lsa, hech qanday harakat kerak emas, javob 0.
    }

    // 1-QADAM: `'a'`, `'b'`, va `'c'` harflarining umumiy sonini sanash
    count := map[byte]int{'a': 0, 'b': 0, 'c': 0} // Har bir harf uchun hisoblagichni 0 dan boshlaymiz
    for i := 0; i < n; i++ { 
        count[s[i]]++ // Har bir belgini hisoblaymiz
    }

    // 2-QADAM: Yaroqlilikni tekshirish (har bir harf kamida `k` marta bo'lishi kerak)
    if count['a'] < k || count['b'] < k || count['c'] < k {
        return -1 // Agar biror harf `k` dan kam bo'lsa, yechim yo'q.
    }

    // 3-QADAM: Eng uzun to'g'ri substringni topish uchun slayd oynasi (sliding window) usulidan foydalanish
    required := map[byte]int{
        'a': count['a'] - k, // Bizga kerak bo'lmagan `'a'` harflar soni
        'b': count['b'] - k, // Bizga kerak bo'lmagan `'b'` harflar soni
        'c': count['c'] - k, // Bizga kerak bo'lmagan `'c'` harflar soni
    }
    maxWindow := 0 // Eng uzun substring o'lchami (boshlang'ich qiymati 0)
    left := 0      // Slayd oynasining chap ko'rsatkichi
    windowCount := map[byte]int{'a': 0, 'b': 0, 'c': 0} // Joriy substringda harflarni hisoblash

    for right := 0; right < n; right++ {
        windowCount[s[right]]++ // Oynaga o'ng tomondan yangi belgi qo'shiladi

        // Oynadagi harflar soni `required` dan oshib ketmasligini ta'minlash
        for windowCount['a'] > required['a'] || 
            windowCount['b'] > required['b'] || 
            windowCount['c'] > required['c'] {
            windowCount[s[left]]-- // Oynadan chap tomondagi belgini olib tashlash
            left++ // Chap ko'rsatkichni bir pozitsiya o'ngga surish
        }

        // Joriy oynaning uzunligini yangilash (agar bu eng uzun bo'lsa)
        maxWindow = max(maxWindow, right-left+1)
    }

    // 4-QADAM: Umumiy uzunlikdan eng uzun substringni ayirish
    return n - maxWindow // Minimal operatsiyalar soni = umumiy uzunlik - eng uzun substring
}

// Yordamchi funksiya: 2 ta sonning kattasini qaytaradi
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

