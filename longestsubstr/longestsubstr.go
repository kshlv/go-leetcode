// Package longestsubstr implements a solution to a problem of
// finding the longest substring where each character is unique
//
// Original problem:
// https://leetcode.com/problems/longest-substring-without-repeating-characters
package longestsubstr

// Status: WIP

func uniqueSubsOfLength(s string, l int) []string {
	if l > len(s) {
		return []string{s}
	}
	subs := subsOfLength(s, l)
	// us := map[string]int{}
	us := []string{}
	for _, sub := range subs {
		m := map[rune]bool{}
		for _, c := range sub {
			m[c] = true
		}
		if len(m) == l {
			us = append(us, sub)
		}
	}

	return us
}

func subsOfLength(s string, l int) []string {
	if l > len(s) {
		return []string{s}
	}
	if l == len(s) {
		return []string{s}
	}
	subs := []string{}
	for i := 0; i+l-1 < len(s); i++ {
		subs = append(subs, s[i:i+l])
	}
	return subs
}

// Substr is actually an extention of
// lengthOfLongestSubstring
func Substr(s string) string {
	allSubs := make([][]string, len(s)+1)
	if s == "" {
		return ""
	}
	for i := range allSubs {
		if i == len(allSubs)-1 {
			sub := allSubs[i]
			return sub[len(sub)-1]
		}
		uniques := uniqueSubsOfLength(s, i+1)
		if uniques == nil || len(uniques) == 0 {
			var subs []string
			if allSubs[i] != nil {
				subs = allSubs[i]
			} else {
				subs = uniqueSubsOfLength(s, i)
			}
			return subs[len(subs)-1]
		}
		allSubs[i+1] = make([]string, len(uniques))
		for j, u := range uniques {
			allSubs[i+1][j] = u
		}

	}
	return s
}

func lengthOfLongestSubstring(s string) int {
	allSubs := make([][]string, len(s)+1)
	if s == "" {
		return 0
	}
	for i := range allSubs {
		if i == len(s) {
			return i
		}
		uniques := uniqueSubsOfLength(s, i+1)
		if uniques == nil || len(uniques) == 0 {
			var subs int
			if allSubs[i] != nil {
				subs = i
			} else {
				subs = i - 1
			}
			return subs
		}
		allSubs[i+1] = make([]string, len(uniques))
		for j, u := range uniques {
			allSubs[i+1][j] = u
		}
	}
	return 0
}
