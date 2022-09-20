package hero

//
//
//                                                  ,/@/@
//                                                 @`   @                    =@]]`             ,]/@@@@@@@@@/
//                                                 O    @                     =@OOO@@@@]@@@@@@@@@@@@@@@@@/
//                                                 O  @@@        ]]]`           \@OOOOOO@@@@@@@@@@@@@@@@\`
//                                                /@   [     /@/[[  [@@\]]        \@OOOOO@@@@@@@@@@@@@@@@@@`
//                                               @@/                             ]@@OOOOO@@@@@@@@@@@@@@@@@`
//                                                                            @@@OOOOOOOO@@@@@@@@@@@@@@@@
//                                                                               ,[[[[[\@@@@@@@@@@@@@@@@@@`
//                                                                                        ,\@@@@@@@@@@@@@@/`
//                                                                                           \O[@@@@@@@@^
//                                                                                            =@   \@@@@@\
//                                                                              ]]]]]@@@@@@@@@@@@@@@\]]`
//                                                                    ]]/@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@]]]
//                                                              ,]@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@]`
//                                     =`                  ,/@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@]`
//                                      \` ]/           ,/@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@]
//                                      /@`           /@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@`
//                                    /[  \        ,@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\
//                            O`    //           ,@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\
//                             ,@`//            /@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@^
//                               @@`          ,@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@`
//                             ,@            =@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\
//                            =/            /@@@@@@@@@@@@@@O@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//                           =^            /@@@@@@@@@@OO@OOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OoOOOO@@@@OO@@@@@@@@@@
//                     [O@@@@@@@^         =@@@@@@OO@@OOO@@@@@@@@@@@@@@@@@@@@@@@@\@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOO@@@OOO@@@@@@@@@
//                         ,/            /@@@@@@@OOOO@@@@@@@@@@@@@@@@@@@@@@@@@@  @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OO@OO@@@@@OO@@@@@@@@@
//                         /            =@@@@@@@@@@@@@@@@@@@@@@@@@@/ =@@@@@@@@^  =@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@O@@@@@@@@^
//                        ,^            @@@@@@@@@@@@@@@@@@@@@@@@@@/  =@@@@@@@/    @@@@@@@@O@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@^
//                        =^           =@@@@@@@@@@@@@@@@@@@@@@@@@@    @@@@@@@     =@@@@@@@^,@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//                        O           ,@@@@@@@@@@@@@@@@@@@@@@@@@@@O   @@@@@@^      \@@@@@@^ ,@@@@@@@@@\/@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@^
//                                    @@@@@@@@@@@@@@@@@@@@@@@@@@`     =@@@@@        @@@@@@^   @@@@@@@@@ \@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//                                   =@@@@@@@@@@@@@@@@@@@@@@@@@^       @@@@^         @@@@@^    \@@@@@@@  =@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//                                   =@@@@@@@@@@@@@@@@@@@@@@@@/        =@@@          ,@@@@@     ,@@@@@@   =@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@^
//                                   @@@@@@@@@@@@@@@@@@@@@@@@@`         =@@           ,@@@@       \@@@@^   =@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@^
//                                   @@@@@@@@@@@@@@@@@@@@@@@@`      ,O\  @/             @@@        ,@@@^    =@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@^
//                                   @@@@@@@@@@@@@@@@@@@ @@@@             `              \@^  ,[      \\     =@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@^
//                                   @@@@@@@@@@@@@@@@@@^  @@\/@@@@@@@@\                       /@@@@@@@@@]     ,@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@^
//                                   @@@@/@@@@@@@@@@@@^    @@/`       \@^                  ,@/[        ,[@@\   ,@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@^
//                                   @@@@ @@@@@@@@@@@@   =@@    /@@@@@\`\\                =/     ]]O@]]    \@\   \@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//                                   =@@^ =@@@@@@@@@@@  =@/   /@@@@@@@@@^\^                    @@@@@@@@@@`   \@`  =@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//                                    \@` =@@@@@@@@@@^  /@`  /@@@@@@@@@@@^                    @@@@@@@@@@@@^   =@`  ,@@@@@@@@@@@@@@@@@@@@@@@@@@@@@^
//                                     @   @@@@@@@@@@   @/   @@@@@@@@@@@@@                   /@@@@@@@@@@@@@`   =@   =@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//                                         =@@@@@@@@@^  @@   @@@@@@@@@@@@@                   @@@@@@@@@@@@@@@    @^  =@@@@@@@@@@@@@@@@@@@@@@@@@@@@^
//                                          \@@@@@@@@@  ,@   @@@@@@@@@@@@@^                 =@@@@@@@@@@@@@@@    @^  =@@@@@@@@@@@@@@@@@@@@@@@@@@@@^
//                                          =@@@@@@@@@`  =^  @@@@@@@@@@@@@^                  @@@@@@@@@@@@@@@    @   =@@@@@@@@@@@@@@@@@@@@@@@@@@@@^
//                                          /@@@@@@@@@^      @@@@@@@@@@@@@^                  @@@@@@@@@@@@@@^   =^   =@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//                                          @@@@@@@@@@\      ,@@@@@@@@`                      @@@@@@@@@@/`[[         =@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//                                          @@@@@@@@@@@        \@@@@@@@@@                    =@@@@@@@@@\            =@@@@@@@@@@@@@@@@@@@@@@@@@@@/
//                                         =@@@@@@@@@@@          \@@@@@/                      =@@@@@@@@@@`          @@@@@@@@@@@@@@@@@@@@@@@@@@@@^
//                                         =@@@@@@@@@@@                                         [@@@@[[             @@@@@@@@@@@@@@@@@@@@@@@@@@@@
//                                         =@@@@@@@@@@/  *oOOO\                                            ,O]]/\* ,@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//                                         =@@@@@@@@@@   [[[O/[                                           *o[`*\[  =@@@@@@@@@@@@@@@@@@@@@@@@@@@^
//                                         =@@@@@@@@@@`                                                            =@@@@@@@@@@@@@@@@@@@@@@@@@@@
//                                         =@@@@@@@@@@@                                                            /@@@@@@@@@@@@@@@@@@@@@@@@@@^
//                                          @@@@@@@@@@@@\                                                          @@@@@@@@@@@@@@@@@@@@@@@@@@@
//                                          @@@@@@@@@@@@@@`                                                       =@@@@@@@@@@@@@@@@@@@@@@@@@@^
//                                          ,@@@^@@@@@@@@@@@@`                 ,\@@@@/[                           =@@@@@@@@@@@@@@@@@@@@@@@@@@
//                                           ,@@ ,@@@@@@@@@@@@@@]                                               ,/@@@@@@@@@@@@@@@@/@@@@^=@@@`
//                                            ,@  ,@@@@@@@@@@@@@@@@@\]                                     ]@@@@@@@@@@@@@@@@@@@@@@ \@@/ ,@@^
//                                                 ,@@@@@@@@@@@@@@@@@@@@@@@@]]]]                ,]]@@@@@@@@@@@@@@@@@@@@@@@@@@@@@/   @/   @^
//                                                   \@@@@@/\@@@@@@@/\@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@`     *@@@@@@@@@@@@@@@@@@@@@@@@@\]
//                                                     \@@@^ ,@@@`[@@  [@@@@@@/  *@@@@@O[`   *  ** /@\*     =@@@@@@@@@@@@@@@@@^******,[@@`
//                                                       ,@^    [@^     /@@@@    O@@[*****  *****  \@/      @@OO@@@@@@@@@@@@[***********,@@`
//                                                                    /@@O@@     @/   ****** **** *//      =@o[O@@@@@OO@@@@`*,\`/O\/O\***,o@\
//                                                                  ,@@OoOO@     =O  *         ]/@/       /@o\oO@OO\,OO\OO@@O@/`    =@@@`**\@^
//                                                                 /@Ooooo/@^     \\*,]]/@@O[[         ,/@`=oooooo[]OOoooOO@@        =@O@\**\@
//                                                                /@Ooooooo/O@` ,\`                 ]@@O\//ooooooo/O@ooooooO@\        \@@@O*=@^
//                                                               @@oooooooo//,@@`   ,`         ]/@@Oooo]oooooooooooOOoooooooO@\       =@@@@^*@^
//                                                              @@\/ooooooooooo`[\@@@@@@@@@@@O[\o/ooooooooooooooo*=@oooooooooO@^      @@O@@^*@@
//                                                             /@ooooOO/ooooooooooo]oo]]/oo\]oooooooooooooooooooo[OOoooooooooo@@`    @@@@@O*=@^
//                                                            =@oooooO@@Oooooooooooooooooooooooooooooo]/OO\/OOOo[`OOooooooooooO@\ ]@@@@@@`**@@
//                                                            @O*ooooOO@OoooooooooooooooooooooooooooooO  ]  ] ,O^/OOooooooooo\\O@@@@@@O`**=@@
//                                                           /@o\ooooo@@oooooooooooooooooooooooooooooo@ OO^=@^ @^=Oooooooooooooo@@^,****/@@^
//                                                          ,@OooooooO@@/oooooooooooooooooooooooooooo/O^ [    ,O^=ooooooooooooooO@@]]O@@@/[`=@
//                                                          /@oooooooO@O^oooooooooooooooooooooooooooooo@^=O@^@^=o^OOooooooooooooo@@O/[,******O^
//                                                         =@ooooooooO@^=oooooooooooooooooooooooooooooooOOoOOoooo^O@oooooooooooooO@O`]OO@@@@O
//                                                         @O*ooO@OOOO@OooooooooooooooooooooooooooooooooooooooooooO@o/ooooo\oooooo@@[O`,[***,@\
//                                                        /@o^*oO@@O@@@ooooooooooooooooooooooooooooooooooooooooooo\@oOoooooOooooooO@@O@@@@@`*=@^
//                                                       /@oooo/oO@OO@@ooooooooooooooooooooooooooooooooooooooooooo/@O@OO@@OooooooooO@    @@@@*=@
//                                                      /@oooooooOOOO@@ooooooooooooooooooooooooooooooooooooooooooooOOOOOoooooooooo=O@^   @@@@O,@^
//                                                     /@\/ooooooooOO@@ooooooooooooo\]oo`\o/\oO]OOOooooooooooooooo^O@^oooooooooooooo@@` =@@@@O`@^
//                                                    //o\oooooooooOO@O^oooooooooO@@@@@@@@@@@@OOOOOooooooooooooooooO@oooooooooooooooO@\,@@@@@/*@
//                                                   @O*=ooooooooooOO@@^/oooooooooooooooooOOOOOOOOOOOoooooooooooo/oO@o,oooooooooooooo@@@@@@@`,O@
//                                                  @@`*oooooooooooOO@@@ooooooooooooooooooooooOOO@@@OooooooooooooooO@O=oooooooooOooooO@@@@/**O@
//                                                 /@@**oooooooooooOO@Oooooooooooooooooooooooooo[ooooooooooooooooooO@@o\oooooooOOoooooO@^*,/@/
//                                               /@` ,@\oooooooooooO@@oooooooooo[[oooooooooooooooooooooooooooooooooOO@@,ooooo/O@OO@O\/o@@@@`
//                                              //     \@OoooooooooO@o*oooooO@@@@OOOOOOoooooooo]oooooooooooooooooooOO@@oooooooO@O@ooooo\@\
//                                              @        ,\@@\]]`=O@Ooooo[`O@o\ooooooooOOOOOOO@@@OOOOoO@O\oooooooooOOO@^=oooooOOOoooooooo@\
//                                             ,@             [[[@@@OoooooO@^oooooooooooooooooooooooooo/O@\oooooooOOOO@O^ooooooooooooooo`=@^
//                                              ,@`        ,       @/,oo/O@o*oooooooooooooooooooooooooooo[\@OoooooOOOO@O/ooooooooooooooooO@@\
//                                               =/@\]             @\/oo^/@^oooooooooooooooooooooooooooooo]/@OoooOOOOO@@^[[\oo[[[\\,]/@@/`  =^
//                                               O  =@,O@@@]`     =@=ooo*O@`ooooooooooooooo/\ooOOOOOOoooooo/O@oooOOOOO@@@@@@@@@@@@/[[`       @
//                                              =^ @@@@/[\@\,/[[[\@Ooooooo@oooooooo/OOOO@@@@@@OOO@@O\oooooooO@/\ooooOO@@                     =^
//                ,@O@`                        ]]@`,@`     =/     =OooooooOOoooooooOOOOOOOOOoooOOoooooooooooo@\oooooOO@@`   ]]]]/O/          =@
//                @OOO@^ ]O@`              ,` =o@` @\]]]/@/ @     =Oooooooooo/oooooooooooooooooooooooooooooooO@oooooOOO@^               ,]@@[
//             ]/@@@@@O@@OO@^             @^ \[O^ =^     /  O     @\ooooooooOOOOOoooooooooooooooooooooooooooooooooooOOO@@@]]]]]]]]/@@@@/
//             @@@@@@@@@@@O@            =@@@/O/O  =^     @ =^    =@[[o\oooooooo[\OOOO@@@@@@@@@@@@@@@@OOOooooooooooooOOOO@^ ,^ @   ,OO^
//             =@@@@@@@@@@@/            @@@@^ =^  @@^    ^  \^    ,@O`,[/ooooooooooooooooooooooooooooooooooooooooOOOOOOO@O    ,\/`
// @@@@@@@@` ]]@@@@@@@@@@@@@`           @@@@^ @   = ,@@@@@/,`@      ,@@OOOoooooooooooooooooooo\*/[[[\oooooooooOOOOOOOOOO@^
//  \@@@@@@@@@@@@@@@@@@@@@@@@@          @@@@^=^ =^  ,]/@/    =\       @@@@@@OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO@@@@@^
//   \@@@@@@@@@@@@@@@@,@@@@@@@\         @@@@\/ O,^      ,^    @       @@@@@@@@@@@@@@@OOOOOOOOOOOOOOOOOOOOOOO@@@@@@@@@@@@^
//    \@@@@@@@@@@@@@@^ \\ @@@@@\]@@@@@@@@@@@@  O O     O`     =^      @OOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@^
//    =@@@@^=@/\@@@@@@`   =@@@@@@@@@@@@@@@@@^  =         ]]]\  ,\     @OOOOOOOOOOOOO@@@@@@@@\ @@O@@@@@@@@@@@@@@@OOOOO@@@^
//    =@@@@  =\,@@@@@@@@\,@@@@@@@@@@@@@@@@@@         ,O[* **,\  =^   ,@OOOOOOOOOOOOOOOOOO@@@^  \@OOOOOOOOOOOOOOOOOOOOO@@\
//     @@@@@`  =@@@@@@@@@@@@@@@@@@@@@@@@@@@@         =^ **=O*\  ,O   =@OOOOOOOOOOOOOOOOOO@@@^   \@OOOOOOOOOOOOOOOOOOOO@@@
//     @@@@@@@/@@@@@@@@@@@@@@@@@@@@@@@@@@@@@          O*,O/* ,^ ,O   =@OOOOOOOOOOOOOOOOOO@@@     @@OOOOOOOOOOOOOOOOOOO@@@
//     @@@@@@@@@@@@@@@@@@@@@@@@OO@@@@@@@@@@@^         O* **]O/  =^   =@OOOOOOOOOOOOOOOOOOO@@      @OOOOOOOOOOOOOOOOOOOO@@^
//     \@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@ @@@`        oO@[`     /    =@OOOOOOOOOOOOOOOOOOO@^       @OOOOOOOOOOOOOOOOOOO@@^
//       ,@@@@@@@@@@@@@@@@OO@@@@@@@@@@@@@`=@@@                 O    /@OOOOOOOOOOOOOOOOOO@O@`       ,@OOOOOOOOOOOOOOOOOO@@^
//             @@@@@O@@@O@@@@@@@@@@@@@@@@ @@@@@\            ]@`    =@@@OOOOOOOOOOOOOOOOO@O@^        =@OOOOOOOOO@@OOOOOO@@@`
//             \@@@@@@@@@@@@@@@@@@@@@@@@/ @@@@/ ,@]    ]/@[        @^*\@@@OOOOOOOOOO@@@@@/O@^        @@OOOOOOOOOOOOOOO@@@\O^
//                 @@@@@@@@@@@@@@@@@@^    ,[[                     =@*   ***,[[[[[[`* *  *=O@^        @@[@O@@@@@@@@/[[****@@
//                 @@@@@@      =@@@@@^                            =@ * ****************  =@@`         @^        ***** *]]@@^
//                 @@@@@^      =@@@@@^                              \\`   ********** *]/@@[            @@@@O]OOOO@@@@@@@@@`O
//                =@@@@@^      =@@@@@                              @@@@@@\]`*]]]/@@@@@@O@^             @[     ,[\@@OOOO@/,@
//                 \@@/         @@@@/                             /@/[[[[[@@@@@@@OOOOO@@@              @^           /@` /`
//                                                               ,^           ,@OOO@@[ ]@              O`[O@@@@@O[`  ,/`
//                                                               /@\]]]]]]]]]/@@/`  ,@`                 ,]]`  ,]]/@[`
//                                                               =^             ,//`
//                                                                 ,[[\@@/[[[[`

type SilentHouseRobot struct {
	heroInfo
}

func (f *SilentHouseRobot) Init() {
	f.Id = 2
	f.Rarity = SSRRarity
	f.MoneyLimit = 100
	f.CurrentMoney = f.MoneyLimit
	f.Name = "静宅机器人"
	f.Motto = "成功计算最优解！静宅机器人是最强的！"
	f.SkillInfo = []SkillInfo{{
		Name: "最优策略",
		Desc: "每当捐赠金币数比本轮最低玩家大1时\n获得4金币"},
	}
}
