/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   printreversealphabet.go                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: yushsato <yushsato@student.42tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/25 18:22:52 by yushsato          #+#    #+#             */
/*   Updated: 2024/06/25 19:21:36 by yushsato         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import "ft"

func PrintReverseAlphabet() {
	for c := 'z'; c >= 'a'; c-- {
		ft.PrintRune(c)
	}
	c := '\n'
	ft.PrintRune(c)
}
