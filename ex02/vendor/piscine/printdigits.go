/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ex02.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: yushsato <yushsato@student.42tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/25 18:22:52 by yushsato          #+#    #+#             */
/*   Updated: 2024/06/25 19:22:05 by yushsato         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import "ft"

func PrintDigits() {
	for c := '0'; c <= '9'; c++ {
		ft.PrintRune(c)
	}
	c := '\n'
	ft.PrintRune(c)
}
