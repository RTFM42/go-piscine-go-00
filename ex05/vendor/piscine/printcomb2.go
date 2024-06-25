/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ex04.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: yushsato <yushsato@student.42tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/25 18:22:52 by yushsato          #+#    #+#             */
/*   Updated: 2024/06/25 18:33:16 by yushsato         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import "ft"

func PrintComb2() {
	for i := 0; i < 100; i++ {
		for j := i + 1; j < 100; j++ {
			ft.PrintRune(rune(i/10) + '0')
			ft.PrintRune(rune(i%10) + '0')
			ft.PrintRune(' ')
			ft.PrintRune(rune(j/10) + '0')
			ft.PrintRune(rune(j%10) + '0')
			if i != 98 || j != 99 {
				ft.PrintRune(',')
				ft.PrintRune(' ')
			}
		}
	}
	ft.PrintRune('\n')
}
