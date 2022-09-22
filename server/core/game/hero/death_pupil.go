package hero

//                                                                                                                          ,`
//                                                                                                                         =O@`
//                                                                                                                        .OOOO@\
//                                                                                                                        ,OOOOOO@`
//                                                                                                                        =OOOOOOOO@`
//                                                                                                                       .OOOOOOOOOOO\
//                                                                                                                       .OOOOOOOOOOOOO.
//                                                                          ...]]]]]]]]]]]]]]`..                         =OOOOOOOOOOOOO^
//                                                               ..]]/OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO\].                 /OOOOOOOOOOO@`
//                                                          .]OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO].           .OOOOOOOOOOOO.
//                                                      ,/@OOOOOOOOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOOOOO].       =OOOOOOOOOO/.        .
//                                                  .]@OOOOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOO@].    =OOOOOOOOO/       ,/`
//                                               ./@OOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOO` .OOOOOOOOO^     ,/@`
//                                            .]@OOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOOOOOO`   ./OO`
//                                          ,/OOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOO@OOOO@` ./OOO`
//                ...]]]]]`..             ,@OOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOO@OO./@OO/.
//         .,/OOOOOOOOOOOOOOOOOOOO\].   ,@OOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOO@OOO/.
//      ,/@OOOOO//[[[[[[OOOOOOOOOOOOOOO@@OOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOO.
//   .`.                       .]/@@@@OOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@O@@@@@@@@@@@@@OOOOO.
//                       .]O@OOOOOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOO@@@@@@@@@@@OOOOO.
//                   ,/OOOOOOOOOOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OO@@@@@@@@@@@@@@OOOOOOO@@@@@@@@@OOOO.
//                ,OOOOOOOOOOOOOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOO@@@@@@@@@OOOOOOOOO@@@@@@@@OOO^
//             ./OOOOOOOOOOOO@@@@OOOO@@@@@@@@@@@@@@@@O@@@@@@@@@@@@@@@@@@@@@@OO@@@@@@@@@@@@@@@@@@@@@@@@@OOO@@@@@OOOO@@OOOOOO@@@@@OOOO.
//           .OOOOOOOOOOOO@@@@@@OOOO@@@@@@@@@@@@@@@@@OO@@@@@@@@@@@@@@@@@@@@@@@@@@OOO@@@@@@@@@@@@@@@@@@@@@@OOOO@OOOO@@@@@OOOOO@@@@OOO^
//         ./OOOOOOOOOO@@@@@@@@OOOO@@@@@@@@@@@@@@@@@@OOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OO@@@@@@@@@@@@@@@@@@@@@@OOOOOOO@@@@@@OOOOO@@@OOO
//        ,@OOOOOOOO@@@@@@@@OOOOOO@@@@@@@@@@@@@@@@@@@OOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOO@@@@@@@@OOOOO@OOO.
//       ,OOOOOOOO@@@@@@`/OOOOOOO@@@@@@@@@@@@@@@@@@@@OOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOO.
//      =OOOOOOO@@@@@` ,@OOOOOOO@@@@@@@@@@@@@@@O@@@@@OOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOO.
//     ,OOOOOO@@@@/.  /OOOO@@OOO@@@@@@@@@@@@@@/=@@@@@@OOO@@O@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOO.
//    .OOOOO@@@@`   .OOOOO@@@OO@@@@@@@@/=@@@@/^=@@@@@@OOOO@@OO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OO.
//    =OOOO@@@`    .OOOOO@@@@OO@@@@@@@^,@@@@O*.=@@O@@@@OOO@@@@O@@@@@@@@@@@@@@@@@@@OO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OO
//   .OOO@@@`     .OOOO@//@/OO@@@@@@@`.=@@@O...O@@*@@@@@OO@@@@@OO@@@@@@@@@@@@@@@@@@@@O][O@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@O^
//   =OOO@`       /OOO@/,@@^OO@@@@@@/..@@@@`...O@O^,@@@@@OO@@@@@@O\\@@@@@@@OO@@@@@@@@OO@@O],\@@@@@@@@@\,\@@@@@@@@@@@@@@@@@@@@@@@@@@@`
//   =OO/        ,OOO@/.=@@.=O@@@@@@^]O@@@/\@\.O@O^.,@@@@@O@@@@@@@@^=@@@@@@O\\\@@@@@@@@\.,[\@\/,\@@@@@@@\..\@@@@@@@@@@@@@@@@@@@@@@@O.
//   \/.         =OO@@..@@@.,@@@@@/@O.=@@@...,@@@O*..,@@@@@O@@@@@@@@//@@@@@@.,O]/@@@@@@@@@`.....[OO@@@@@@@\..,\@@@@@@@@@@@@@@@@@@@@^
//  ,^           =OO@^ .@@^  \@@@@[=@^=@@O.....=@@@`..,@@@@@@@@@@@@@@.[@@@@@^..\`*,\@@@@@@@@@`....,]]/@@@O@@\]]`\@@@@@@@@@@@@@@@@@O.
//               =O@@` .@@^  ,@@@@\]/@O@@@OO\]]]@@O@@\..O@@@@@@@@@@@@@.,@@@@^....\`....[O@@@@@@@@[...*.,@\,o`\@`.,@@@@@@@@@@@@@@@@`
//               =O@@. .@@^   =@@@/[[[\@@O,[[[[[\@/[\@@@@@@@@@@@@@@@@@@/@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\`.\..\`.=@@@@@@@@@@@@@^
//               .O@O  .@@^   =@@@@`....@@@@/[[[\@@@`.\@@/@@@@@@@@@@@@@@/O@@^......,]@@@@/\]@@@@@@@@O]`..[@@^.......,@@@@@@[O@@@^
//                =@^   =@^  .@/\@@@`.../@`..../O@@@@@\*\@`@@,@@@@@@@@@@@O@@^,`]/@@@@/`]@@/,`....,@@^,@@@O`=@^....]=@@@@`.......\.
//                 =^   .@\]/@@@@@@@@\,@/.....=^...@@^.\^\`,@^..,\@@@@@@@@@O.O/O[\/\//`..../......,@.,@@.,\@@\.]]]O@@@`...,]OO`.=^
//                  =.   =@@@@@[[\/@@@\@^......^....@`.=@\,=@@@@@@@@@@@@@@@@@`.[[.[`...............^.@@@^..@\@@@@@@@/...//]\^...=^
//                   .    \,@@^   ,`\@@\O...........=..@@O.,\@@OOOOO@@@@./^.[[\`..................=`=@@@....=@@/......,O/oo^O...=.
//                        .^ =^    =`.\@\......,....=.=@@^...O........=^.O......[..........,\`..,/`.@@@^....=@......./OOO/\\/...^
//                         =`.^     \...\\......\..,@\O@@`...\.........^....................,OoOOOOOoO^.....=O.....=O//ooo^O...=.
//                           .^     .O...,^......,OOOOOO`....,.........O............,\].......,[OOO/`]/.....O`....=O/oooo//`..,^
//                            =.     .\...=`.,\.......,/`..............,\................[\\]]`.]]O[.......=`.....=^oooo//`.../.
//                            .\      ,^...\....[\oO[`.=`...............,\................................/`....../oo/\//...,/.
//                             .\      ,^...\...........,@]]`.............\`............................,/........OO/`....,/.
//                              .\`     .\`.=^............\\\/O]............\]........................]/.....]O[`.......][.
//                                .[].    .\]\.........,/...\\\//..............[O]].............]]//[.....,/.........,/`
//                                    .[\].. ,\....]//`................................................../`......../[.
//                                             ,\...............................]]O/........................]]/@@@@@`
//                                               ,]................]]]]/OO[[[........................,/[[\O@@@@@@@@@@@@`
//                                                 ,\..............................................,/.            ..,[[[OO`
//                                                   .[\........................................]O[=^
//                                                      .[\.................................]O/\\/O\=@\.
//                                                         .[\..........................]O/\\/O/]oo\@@O@.
//                                                            .[\`.................,/O[`\/O[]ooooo/@OOOO@`
//                                                                ,\].........,/O[\]/O/\]ooooooo//@OOOOOO@].
//                                                                  ..]@@@//[\/[[\]/[[ooo[[[\o/./@OOOOOOO@OOO@@@\]`.
//                                                         ..]]O@@@@@@@@@@^o//Oo\@]@OOOOOOOO^..OOOOOOOOO@@OOOOOOOOOO@@@O].
//                                                    ,/@@@OOOOO@OOOOOOOOO@\OOOO@OOOOOOOOOO@.,@OOOOOOOOO@OOOOOOOOOOOOO@@@@\.
//                                                  ./@@@OOOOOOOO@OOOOOOOO@OOOOO@OOOOOOOOOO^=@OOOOOOOOO@@OOOOOOOOO@@@OOOOOO@^
//                                                 .@OO@OOOOOOO@@@@OOOO@@@@OOOO@@@@@@OOOO@^/@OOOOOOOOOO@OOOOOOOOO@OOOOOOOOOOO@`
//                                                ,@OO@@OOOOOO@@@OOOOOOO@@@@OOOO@@@@OOOOOO@@OOOOOOOOOO@@OOOOOOOOO@OOOOOOOOOOOO@\.
//                                               =@OOO@OOOOO@@O@@OOOOOOOOO/o,O[/=@OOOOOOOOO@OOOOOOOOOO@OOOOOOOOO@@OOOOOOOOOOOOOO@^
//                                             ./@OOO@@OOOO@@OO@OOOOOO@@/.=`..=..\OOOOOOOO@@OOOOOOOOO@@OOOOOOOOO@OOOOOOOOOOOOOOOOO@`
//                                            .OOOOOO@OOOOO@@@@@OO@@OOO.../...^.,/...../@O@@OOOOO@@@@OOOOOOOOOO@@OOOOOOOOOOOOOOOOOO@\
//                                           .@OOOOO@@OOOOOOO@@@@@OOO@`..=.../\/......=@OOO@O@@@@@@@OOOOOOOOOOO@@OOOOOOOOOOOOOOOOOOO@^
//                                          ,@OOOO@@@@OOOO@@@OOOOOOO@/...^..=/.......=@OOOOOO@@@@@OOOOOOOOOOOOO@OOOOOOOOOOOOOOOOOOOO@.
//                                         .@OOO@@@@@OOOOO@OOOOOOOOO@.../...^.......=@OOOOOOOOOOOO@@OOOOOOOOOO@@OOOOOOOOOOOOOOOOOOO@^
//                                         .OOOO@@@@@OOOOO@OOOOOOOO@^..,^../.......=@OOOOOOOOOOOO@@OOOOOOOOOOO@@OOOOOOOOOOOOOOOOOOO@O@`
//                                          =@OO@@@@OOOOOO@@OOOOOOO@.../..=`....../@OOOOOOOOOOOO@@OOOOOOOOOOOO@@OOOOOOOOOOOOOOOO@@@@OOOO.
//                                          .@OO@@@@OOOOOO@@OOOOOOO/..=`..^...../@OOOOOOOOOOOO@@OOOOOOOOOOOOOO@OOOO@@@@@@@@@@@@@@@@@OOO@^
//                                           =@OO@@OOOOOOO@@OOOOOOOO..^../..../@OOOOOOOOOOOOO@@OOOOOOO@@OOOOOO@OOO@@OOOOO@@@@@@@@@@@@O@^
//                                          =@@OO@@OOOOOOOO@@OOOOOO@`=...`../@OOOOOOOOOOOOOO@OOOOOOOOO@@@@@@OO@@OOOO@@@OOOOOOOOO@@@OO@@^
//                                         =@@@OO@@OOOOOOOOO@@OOOOO@O^..../@OOOOOOOOOOOOO@@@OOOOOOOOOO@@@@@@@@@@@@@@OOO@@@OOOOOOOOOO@@O@^
//                                        ,@OO@OO@OOOOOOOOOOO@@OOOOO@=../@OOOOOOOOOOO@@@OOOOOOOOOOOOO@@@@@@@@@@@@@@\@@@OOO@@@OOOOOO@@OOO@\
//                                        \@OO@O@@OOOOOOOOOOOO@@OOOO@\O@OOOOOOOOOO@@@OOOOOOOOOOOOOOOO@@@@@@@@@@@@@^ \@O@@@OOOO@@@OO@OOOOO@\.
//                                         ,@OOO@@OOOOOOOOOOOOO@@OOO@@OOOOOOOO@@@OOOOOOOOOOOOOOOOOOOOO@@@@@@@@@@@/  .\OOOOOOOOOOOOOOOOOOOOO\.
//                                           =@O@@OOOOOOOOOOOOOO@@@@OOOOOO@@@OOOOOOOOOOOOOOOOOOOOOOOOOO@@@@@@@@@@.   .\OOOOOOOOOOOOOOOOOOOOOO.
//                                           =@@@@OOOOOOOOOOOOOO@@OOOOO@@@OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO@@@@@@@^     .@OOOOOO@@@@@OOOOOOOOOO@.
//                                          .OOO@@OOOOOOOOOOOO@@OOO@@@OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO@@@@@O.    =@@@@@@@@@@@@@@@@@@OOOOOO@`
//                                          .@OO@@OOOOOOOOOOO@O@@@OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO@@@@`     .[O@@OOOOOO@@@@@OOOOOOOOO@^
//                                          =@OOO@OOOOOOOOO@@O@OOOO@@@@OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO@^         ./O@@@OOOOOOOOOOOOOOO@^
//                                         .@OOOO@OOOOOOOOOO@@@O@OOOO@OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO@`         ,@OOOOOOOOOOOOOOOOOOO@^
//                                        .@@OOOO@OOOOOOOOOO@@OO***,OO@OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO/.         =@OOOOOOOOOOOOOOOOOOO@`
//                                       ./@@OOOO@OOOOOOOOOO@@@\****Oo@OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO@^         .OOOOOOOOOOOOOOOOOOOOO@`
//                                       /@@OOOO@@OOOOOOOOOOO@O@\**/O@OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO@`         ,@OOOOOOOOOOOOOOOOOOOOO.
//                                      =@OOOOO@O@@OOOOOOOOOO@OOO@@@OOOOOOOOOOOOOOOOOOOO@@@@OOOOOOOOOOO/.         /@OOOOOOOOOOOOOOOOOOOO/.
//                                     .OOOOOO@OO@OOOOOOOOOOO@@OOOOOOOOOOOOOOOOOOOOO@@/`.....,[O@@OOO@@`        .@OOOOOOOOOOOOOOOOOOOO@/
//                                     =@OOOO@@OO@OOOOOOOOOO@@@OOOOOOOOOOOOOOOOOO@@`..../@\`....^..,[O@@O`     =@OOOOOOOOOOOOOOOOOOOO@^
//                                     /OOOO@@OOOOOOOOOOOOOO@O@OOOOOOOOOOOOOOOO@/[\`,/@@@@@@@@/`.........,\\],/.[O@OOOOOOOOOOOOOOOOO@`
//                                    ,@OOO@@OOOOOOOOOOOOOO@@O@OOOOOOOOOOOOOOO@\]/@@OOOOO@O`...................[\..=@OOOOOOOOOOOOOO@`
//                                    =@OOO@OOOOOOOOOOOOOOO@OOOOOOOOOOOOOOOOOOOOOOOOOO@/`.....,/................=\..=@OOOOOOOOOOOO/.
//                                   .@OOO@@OOOOOOOOOOOOOO@@OOOOOOOOOOOOOOOOOOOOOOOO@/.....,//..................*/^..OOOOOOOOOOO@/
//                                   =@OOO@OOOOOOOOOOOOOOO@@OOOOOOOOOOOOOOOOOOOOOOO@`.../@@/..../`...............^O...@OOOOOOOO@`
//                                  .OOOO@@OOOOOOOOOOOOOOO@@OOOOOOOOOOOOOOOOOOOOO@^.../@@/....,/...,`............o=^..=@OOOOOO@.
//                                  =@OOO@@OOOOOOOOOOOOOOO@@OOOOOOOOOOOOOOOOOOOO@]..,@@@^.../@`..,@@@@OOOoooo/[[O[\^.../@OOO@/.

type DeathPupil struct {
	heroInfo
}

func (f *DeathPupil) Init() {
	f.Id = 6
	f.Rarity = SSRRarity
	f.MoneyLimit = 100
	f.CurrentMoney = f.MoneyLimit
	f.Name = "名侦探"
	f.Motto = "凶手有很多，死神只有一个"
	f.SkillInfo = []SkillInfo{{
		Name: "死神",
		Desc: "上回合每有一名玩家被淘汰\n获得3金币"},
	}
}

func (l *DeathPupil) OnRoundStart(ctx *SkillContext) {
	if l.IsBankrupt() {
		return
	}
	for i := int64(0); i < ctx.BankruptcyCount; i++ {
		l.AddMoney(3)
	}
}
