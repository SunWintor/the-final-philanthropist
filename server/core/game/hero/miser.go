package hero

//                                                                   ,@`              ./\
//                                                             ...@@@ ,].       .  @@@..`
//                                                            =@@^@@^@@@`      =@@`@@@@@@^
//                                                             =@@@@@@/        .@@@@@@@@`
//                                                               ,@@@@\          ,@@@@\
//                                                              ,@`***@@.       //****\@.
//                                                             .@^*****@^      =/******@\
//                                                             =@******=@.    .@*******=@.
//                                                             @^******,@^    =^*******=@.
//                                                            .@`*******@^    @^*******=@^
//                                                            .@********@^   .@********=@.
//                                                            =@********@^   =@********=@.
//                                                            =@********@^   =@********@@
//                                                            =@********@^   =@********@^
//            ,]@@@@@@@@@@@].                                 =@********@^   =^*******,@^
//        ,/@@@@/[[\@@@@@@@@@@\.                              .@********@^   =@*******=@.
//      /@@@/*******`\@@@@@@@@@@@.                            .@^*******@^   =@*******@@
//    =@@@/**********=@@@@@@@@@@@@`                            \^**@@@^*@^   =@**/@`**@^
//  .@@@@\*********`/@@@@@@@@@@@@@@`                           =@*/\@\@`@`   =@.@..=^=@.
// .@@@/*********,]@@@@@@@@@@@@@@@@^                           .@`\@@@@*@.   .@=@@@@@/@
// /@@O,********/@@@@@@@@@@@@@/[@@[.                            =^*@@@`=@.   .@*@@@@`@^
//=@@@*******`/@@@@@@@@@@@@@@@@                                 ,@*****=@.   .@^*[[*=@.
//@@@/*****,,@@@@@@@@@@@@@@@@@`                                  \^****=@     @^****@^
//@@@^*****,@@@@@@@@@@@@@@`         ..                           ,@****=@     @^***,@`
//@@@^***,=@@@@@@@@@@@@@@@@`      =@@@@@\`                        \^***=@     =^***=@
//@@@@``,=@@@@@@@@@@@@@@@@@^      =@@@@@@@\.              .       .@***=^     =\***@^
//@@@@@@@@@@@@@@@@@@@@@@@/.     ./@@@@@@@@@@.             =@`      =^**=^ ..  =@**=@.
//=@@@@@@@@@@@@@@@@@@@@@@@@`    =@@@@@@@@@@@^             .\@@@@@@`.@*/@@@@@@@@@**/^
//.@@@@@@@@@@@@@@@@@@@@@@@@/   .]/@@@@@@@@@@@                  ,\@@@@@@@@@@@@@@@**@\.
// ,@@@@@@@@@@@@@@@@@@@@@@/   .@@@@@@@@[\@@@^           ]@@@@@\` =@@@@@@@@@@@@@@@@@@@\]]].
//  =@@@@@@@@@@@@@@@@@@@@@^   ,@@@@@@``*=@@@          ,@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\.
//   ,@@@@@@@@@@@@@@@@@@@@@@@@@@@@@/**`=@@/          .@@@/ .     ,[@@@@@@@@@@@@@[[.    .\@@@@.
//     \@@@@@@@@@@@@@@@@@@@@@@@@@`/*,,@@@`           .@@@^@\`                         =@[=@@@^
//      .\@@@@@@@@@@@@@@@@@@@@@@@/O@@@@`              \@@@   \@\`                ,,]@[ ../@@@^
//         ,\@@@@@@@@@@@@@@@@@@@@@@@`                 .@@@\.     .[\@@@]]]@]/@@/[[^     =@@@/
//             ,[@@@@@@@@@@@@[[.                        \@@@\             ..          ,@@@@@@`
//                    =@@@@@@^                           .@@@@@]                  .]@@@@@@@@@@`
//                     @@@@@@^                          /@@@@@@@@@@@]]`.....]]/@@@@@@@@@@@@@@@@@@@@@@@@\]`
//                     ,@@@@@\                ........]@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@[*,`**,[***,*/@@@]]`.
//                      =@@@@@`             =@\****,*/@,@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@[*,*************@@@@@@@@@@@@@@@@\]]`.
//                       @@@@@@\]]]`....]]]]/\@@\*/=@*@^//@@@@@@@@@@@@@@@@@@@@@@[`**`**********],*,**,=@@@@@@@@@@@@@@@@@@@@^
//                       .\@@@@@@@@@@@@@@@@@@@@@@*@^/*=/,^@^,[\@@[\@\***,`**=,@****************\,\@@\]*,[O@.     .[[\@@@@@@.
//                         ,@@@@@@@@@@@@@@[[[\[@@@,`******@^]@/,,***`[@@`*`,,@*******************,`@\,[[[.          =@@@@@.
//                                        ,/@@@@/,********=/`/*********,,\@/@,**********************\@.            =@@@@@.
//                                       =@OO@@`**************************,*,***********************`\@@@@`       =@@@@@.,]]]]].
//                                       ,@OOOO@\,***************************************************`/@OO@.     /@@@@@@@@@@@@@@@@\`
//                                      ,@@@@OOOOO@\`*,*****************************************`]/@@@OOO@@     /@@@@@[****`O@@@@@@@@`
//                                    ./@@@@@OOOOOOOO@@@\]***`************************`***]]@@@@OOOOOOOOOO@@. .@@@@@^`****]@@@@@@@@@@@^
//                                   ,@@@@@@OO@@@@@@O@@@@OO@@@@@]]]`***********]]]]@@@@@@OOOOOOOOOOO@@@@@OOO@]@@@@@/**/]@@@@@@@@@@@@@@@`
//                                  =@OOO@@OO@@@@@@\@@@@O//@@@OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO@@@@@@@@@@@OOO@@@@@@^`/@@@@@@@@@@@@@@@@@^
//                                 /@OOOO@OO@@@@@@^@@@@@@@@@o@@OOOOOOOOOOOOOOOOOOOOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@/  ]/@\`
//                                =@OOOOOOOOOO@@@@/@@@@@@@@@O@@@@@@@@@@@@@@@@@@@@@@@@OOOOOO@@@@@@@@@@@@@@OO@@@@@@@@@@@@@@@.,[` ,[`     .@@@@@@@\
//                               ,@OOOOOOOOOOOOO@o@@@@@@@@@\@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOO@@@@@@@@@@OOOOO@@@@@@@@@@@@@@@@@`        ./@@@@@@@@@@@.
//                               =@OOOOOOOOOOOOOO@@@//@@@@OO@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOO@@@@OOOOOOOOOOO@@@@@@@@@@@@@@@@@^  .]`   =@@@@@@@@@@@@^
//                               =@OOOOOOOOOOOOOOO@OOO@@@@@@OOOOOO@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOOOOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@`.@@@@@@@@@@@@@@
//                               .@@OOOOOOOOOOOOOO@OOO@@@OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//                                ,@OOOOOOOOOOOOOO@OOO@@@OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@O[[\@@@@@
//                                 ,@OOOOOOOOOOOOO@OOO@@@OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@``***,=@@@@
//                                  .@@OOOOOOOOOOO@OOO@@@OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@/`*******=@@@/
//                                    ,@@OOOOOOOOO@OOO@@@OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@/,`*******`O@@@.
//                                      \@OOOOOOOO@@OOO@@@OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO@@@@@@@@@@@@@@@@@@@@`/********,,@@@@.
//                                        [@@OOOOOOO@OO@O@OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO@@`,@@@@@@@@@@@@@@@[*`*******,*]@@@@`
//                                          ,@@OOOOOOO@@@@OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO@@@`     .\@@@@@@@@@@@,`*******,]@@@@@`
//                                             ,@OOO@@OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO@@@[.           .[@@@@@@@@@@@OO@@@@@@@@[.
//                                             =@OOOO@@@OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO@@@@@/`                    ,[\@@@@@@@@@@[[.
//                                             @OOOOOOOOO@@@OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO@`
//                                             ,@@@@OOOOOOOOO@@@@@@@@OOOOOOO@@@@@@@@OOOOOOOOO@\
//                                                =@@@@@@@@@@@@@   ..,[[[[[[=@OOOOOOOOOOOOOOOO@.
//                                                 @@@@@@@@.                .\@@@@OOOOOOOOOOOO@^
//                                                 .@@@@@@^                      @@@@@@@@@@[..
//                                                  ,@@@@@                       ,@@@@@@@@/
//                                                   .@@/.                        =@@@@@@@`
//                                                                                 \@@@@@^
//                                                                                  \@@@^

type Miser struct {
	heroInfo
}

func (f *Miser) Init() {
	f.Id = 20
	f.Rarity = NRarity
	f.MoneyLimit = 110
	f.CurrentMoney = f.MoneyLimit
	f.Name = "吝啬鬼"
	f.Motto = "我穷得只剩钳了，我捐10金币"
	f.SkillInfo = []SkillInfo{{
		Name: "一掷十金",
		Desc: "捐赠金币数量最多为10"},
	}
}
