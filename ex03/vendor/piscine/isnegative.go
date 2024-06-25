/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ex03.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: yushsato <yushsato@student.42tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/25 18:22:52 by yushsato          #+#    #+#             */
/*   Updated: 2024/06/25 18:29:33 by yushsato         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import "ft"

func IsNegative(nb int) {
	if nb < 0 {
		ft.PrintRune('T')
	} else {
		ft.PrintRune('F')
	}
	c := '\n'
	ft.PrintRune(c)
}
