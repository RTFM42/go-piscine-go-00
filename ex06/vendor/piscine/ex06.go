/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ex06.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: yushsato <yushsato@student.42tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/25 18:22:52 by yushsato          #+#    #+#             */
/*   Updated: 2024/06/25 19:14:24 by yushsato         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import (
	"ft"
)

func combn(lastIndex int, buf []byte, offset *int, index int) {
	for buf[index] <= '9' {
		if index == lastIndex {
			for i := *offset; i <= index; i++ {
				ft.PrintRune(rune(buf[i]))
			}
			if *offset == 2 {
				*offset = 0
			}
		} else if buf[index] == '9' {
			break
		} else {
			buf[index+1] = buf[index] + 1
			combn(lastIndex, buf, offset, index+1)
		}
		buf[index]++
	}
}

func PrintCombN(n int) {
	buf := make([]byte, 11)
	var offset int

	if n != 0 {
		offset = 2
		buf[0] = ','
		buf[1] = ' '
		buf[2] = '0'
		combn(n+1, buf, &offset, 2)
		ft.PrintRune('\n')
	}
}
