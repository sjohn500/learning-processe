# Audit Prep — `Capital` and `CapN` Functions

---

## The Core Idea (say this if asked to summarise)

> "Both functions walk a word slice, find a marker token, remove it cleanly using append reslicing, then transform the target words by splitting each word at its first character to capitalize it."

---

## Function 1 — `Capital`

### What it does
Finds `(cap)` in the text and capitalizes the word immediately before it, then removes `(cap)` from the output.

### The code
```go
func Capital(txt string) string {
    words := strings.Fields(txt)
    for i := 0; i < len(words); i++ {
        if words[i] == "(cap)" && i > 0 {
            words[i-1] = strings.ToUpper(words[i-1][:1]) + strings.ToLower(words[i-1][1:])
            words = append(words[:i], words[i+1:]...)
            i--
        }
    }
    return strings.Join(words, " ")
}
```

### Step by step example
Input: `"hello (cap) world"`

1. `strings.Fields` splits into: `["hello", "(cap)", "world"]`
2. Loop finds `"(cap)"` at index `i = 1`
3. Capitalizes `words[i-1]` → `words[0]` = `"hello"` → `"Hello"`
4. Deletes `"(cap)"` from slice
5. Result: `"Hello world"`

---

## Function 2 — `CapN`

### What it does
Finds `(cap, N)` and capitalizes the **N words before it**, then removes the marker from the output.

### The code
```go
func CapN(va string) string {
    words := strings.Fields(va)
    for i := 0; i < len(words); i++ {
        if words[i] == "(cap," {
            num := strings.TrimSuffix(words[i+1], ")")
            words = append(words[:i], words[i+2:]...)
            i--
            val, err := strconv.Atoi(num)
            if err != nil {
                fmt.Println("convertion fail")
            }
            start := i - val + 1
            if start < 0 {
                start = 0
            }
            for j := start; j <= i; j++ {
                if len(words[j]) > 0 {
                    words[j] = strings.ToUpper(words[j][:1]) + strings.ToLower(words[j][1:])
                }
            }
        }
    }
    return strings.Join(words, " ")
}
```

### Step by step example
Input: `"hello world go (cap, 2)"`

1. `strings.Fields` splits into: `["hello", "world", "go", "(cap,", "2)"]`
2. Loop finds `"(cap,"` at index `i = 3`
3. `words[i+1]` = `"2)"` → `TrimSuffix` removes `)` → `num = "2"`
4. Delete both `"(cap,"` and `"2)"` → slice becomes `["hello", "world", "go"]`
5. `i--` → `i = 2` (pointing at `"go"`)
6. `strconv.Atoi("2")` → `val = 2`
7. `start = 2 - 2 + 1 = 1`
8. Inner loop: capitalize indexes 1 and 2 → `"World"`, `"Go"`
9. Result: `"hello World Go"`

---

## The 7 Questions They Will Ask

---

### Q1. Explain this line:
```go
words = append(words[:i], words[i+1:]...)
```

**Answer:**
> "This removes the element at index `i` from the slice.
> `words[:i]` is everything before index `i`.
> `words[i+1:]` is everything after index `i`.
> `append` joins them together, skipping the element at `i`.
> The `...` unpacks the second slice so append receives individual elements, not a slice of a slice."

**Visual:**
```
["hello", "(cap)", "world"]
           i=1

words[:1]  = ["hello"]
words[2:]  = ["world"]

append(["hello"], "world") = ["hello", "world"]  ✓
```

---

### Q2. Why `i--` after deleting from the slice?

**Answer:**
> "When you delete an element, the slice shrinks by 1. Every element after the deleted one slides left by one position. Without `i--`, the loop's `i++` would skip the element that just moved into position `i`. Doing `i--` corrects for that shift."

**Visual:**
```
Before:  ["hello", "(cap)", "world"]
                    i=1
After delete: ["hello", "world"]
                    i=1 → now points to "world"
Loop does i++ → i=2 → SKIPS "world"!  ✗

With i--: i=0, loop does i++, i=1 → correctly lands on "world"  ✓
```

---

### Q3. What does `[:1]` and `[1:]` mean?

```go
strings.ToUpper(words[i-1][:1]) + strings.ToLower(words[i-1][1:])
```

**Answer:**
> "`[:1]` slices from the beginning up to but NOT including index 1 — so just the first character.
> `[1:]` slices from index 1 to the end — the rest of the word.
> We uppercase the first character, lowercase the rest, then join them."

**Visual:**
```
"hello"
  ↓           ↓
[:1] = "h"   [1:] = "ello"
  ↓           ↓
"H"          "ello"
  └─────┬─────┘
      "Hello"
```

---

### Q4. Why `strings.TrimSuffix(words[i+1], ")")`?

**Answer:**
> "`strings.Fields` splits on spaces, so `(cap, 2)` becomes TWO tokens: `"(cap,"` and `"2)"`.
> The number `2` comes attached to a closing bracket `)`.
> `TrimSuffix` removes the trailing `)` so we are left with just `"2"`, which can then be converted to an integer."

```
"(cap, 2)"  →  Fields  →  ["(cap,"]  ["2)"]
                                        ↓
                               TrimSuffix(")") = "2"
                                        ↓
                               strconv.Atoi("2") = 2
```

---

### Q5. What does `strconv.Atoi` do and what is `err`?

**Answer:**
> "`Atoi` stands for ASCII to integer. It converts a string like `"3"` to the integer `3`.
> It returns two values — the converted number and an error.
> If the string is not a valid number, `err` will not be nil and `val` will default to `0`."

```go
val, err := strconv.Atoi("2")   // val = 2, err = nil       ✓
val, err := strconv.Atoi("abc") // val = 0, err = not nil   ✗
```

**Weakness to acknowledge:**
The current code prints a message on error but continues running. A stronger implementation would return early to avoid incorrect behaviour when `val` is `0`.

---

### Q6. Explain `start := i - val + 1`

**Answer:**
> "After deleting the marker, `i` points to the last word we want to capitalize.
> `val` is how many words to capitalize.
> `start` is the index of the first word in that range.
> The `+1` is there because both `start` and `i` are inclusive — we want exactly `val` words."

**Example:**
```
i = 4, val = 3
start = 4 - 3 + 1 = 2

Capitalize indexes: 2, 3, 4  →  that is exactly 3 words  ✓
```

---

### Q7. Why `if start < 0 { start = 0 }`?

**Answer:**
> "If the user writes `(cap, 10)` but there are only 2 words before it, the math gives a negative index.
> Accessing a negative index in Go causes a panic (crash).
> This guard clamps `start` to `0` so we just capitalize whatever words exist, instead of crashing."

```
i = 1, val = 10
start = 1 - 10 + 1 = -8   ← would panic without the guard

With guard: start = 0  ✓   (capitalizes words 0 and 1 instead)
```

---

## The One Weakness to Know

In `CapN`, if `(cap,` appears at the very end of the input with no number after it:

```go
num := strings.TrimSuffix(words[i+1], ")")
```

`words[i+1]` would be **out of bounds** — this causes a panic. The code does not guard against this case. If the auditor raises it, acknowledge it confidently:

> "You are right, there is no bounds check before accessing `words[i+1]`. A fix would be to check `if i+1 < len(words)` before that line."

---

## Quick Cheat Sheet

| Question | One-line answer |
|---|---|
| What does `append(words[:i], words[i+1:]...)` do? | Removes element at index `i` from the slice |
| Why `i--` after delete? | Corrects for the slice shrinking so no word is skipped |
| What does `[:1]` do? | Takes only the first character |
| What does `[1:]` do? | Takes everything except the first character |
| Why `TrimSuffix`? | Removes the `)` stuck to the number token |
| What is `strconv.Atoi`? | Converts a string number to an integer |
| What does `start := i - val + 1` calculate? | The index of the first word to capitalize |
| Why `if start < 0`? | Guards against a panic when val is larger than available words |

---

*Good luck tomorrow Serah — you've got this!* 🎯
