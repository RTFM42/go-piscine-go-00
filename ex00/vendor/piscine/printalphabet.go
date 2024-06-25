/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   printalphabet.go                                   :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: yushsato <yushsato@student.42tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/25 18:20:22 by yushsato          #+#    #+#             */
/*   Updated: 2024/06/25 19:20:52 by yushsato         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import "ft"

func PrintAlphabet() {
	for c := 'a'; c <= 'z'; c++ {
		ft.PrintRune(c)
	}
	c := '\n'
	ft.PrintRune(c)
}
