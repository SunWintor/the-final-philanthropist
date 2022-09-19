package hero

//
//
//
//
//                                                                                                        [O`
//                                                                                                          \O`
//                                                                     ,]/O@@@@@@@@@@@@@@@@@@@@@@@O]]        \@O
//                                                               ,/O@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@O]`   \@@^
//                                                           ,/@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\  O@@\
//                                                        ,O@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OO@@@O`     /^
//                                                     ]O@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOO@@@@@@@@@@@@@@@O@@@@@@@@@@@`  ,O@`
//                                                  ,O@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOO@@@@@@@@@@@@OO@@@@@@@@@@@@^ O@O
//                                                /@@@@@@@@@@@@@@@@@@@@@OOO@@@@@@@@@@@@@@@O@@@@@@@@@@@@@OO@@@@@@@@@@@@@@@@/
//                                              O@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\
//                                           ,O@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@O
//                                          /@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OO@@@@@@@@@OO@@@@@@@@@@@@@@@@@@@@@@O
//                                        /@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@O[\OO@^,O@OO@O[*/@OO[[ O@@@@@@@@@@@@@@@@@O`
//                                      /@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OO/[[[[[[[[[\O^      [`                 OOO@@@@@@@@@@@@O`
//                                    ,@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@/[`                                        \@@@@@@@@@@@@@@\
//                                   /@@@@@@@@@@@@@@@@@@@@@@@@@@O[                                                O@@@@@@@@@@@@@\
//                                  /@@@@@@@@@@@@OO@@@@@@@@@@/                                                     O@OO@@@@@@@@@@\
//                                ,@@@@@@@@@@@OOOO@@@@@@@@/`                                                        \@^,O@@@@@@@@@\
//                               /@@@@@@@@@@OOO@@@@@@@@O`                                                            O@^ \@@@@@@@@@\
//                              =@@@@@@@@@@@@@@@@@@@@O`                                                              ,OO` =@@@@@@@@@^
//                              O@@@@@@@@@@@@@@@@@@@^                                                                 =@\  =@@@@@@@@@`
//                              O@@@@@@@@@@@@@@@@@@O                                                                   OO`  =@@@@@@@@^
//                             =@@@@@@@@@@@@@@@@@@@^                                                                   =O^   O@@@@@@@@`
//                            ,@@@@@@@@@@@@@@@@@O@O                                                                     \^   =@@@@@@@@O
//                            O@@@@@@@@@@@@@@@@OO@O                                                                     *     O@@@@@@@@`
//                           =@@@@@@@@@@@@@@@@OoO@O                                                                     =^    =@@@@@@@@^
//                           O@@@@@@@@@@@@@@@@O`O@^                                                                     /^     O@@@@@@@O
//                          =@@@@@@@@@@@@@@@@@O*O@^                                                                     [      =@@@@@@@@`
//                          O@@@@@@@@@@@@@@@@@^ O@^                                                                            ,@@@@@@O@^
//                         ,@@@@@@@@@@@@@@@@@O` /@O                                                        ]]]]OOOOOOO/`        \@@@@@O@O
//                         =@@@@@@@@@@@@@@@@@/  =@O            ]]OOO@@@OOO\]`                         ,/@@@@@@@@@O[             =@@@@@/O@
//                         O@@@@@@@@@@@@@@@@O^  ,@O      ,/O@@@@@@@@@@@@@@@@@@`                     /@@@@@@@@OO`   *            =@@@@@^=O
//                        =@@@@@@@@@@@@@@@@@o*   \@`  /OO[`      [[[OO@@@@@@@@@O*                 ,@@@@@/[                      =@@@@@^ O
//                        =@@@@@@@@@@@@@@@@O\*   ,O^                ********                     /@/[*   ******                 =@@@@@^ o^
//                        O@@@@@@@@@@@@@@@O`*     ,O         ]OOOOOOOOOO[[[[**                     ****,/oo[[[[[[`   *[[[*      *oO@@@^ \^
//                        O@@@@@@@@@@@@@@@^ *      =^   ,/o/[`            **`                        ,/`   ]/OO@@@@@@@@@@@@@O    *O@@@^ =
//                        O@@@@@@@@@@@@@@O^ *      ,O/*  ,]]ooOoooOOO]]]]*                            ,/@@@@@@@@@@OO[`            =@@@^
//                        OOO@@@@@@@@@@@O\^      `   ]/O@@@@@@@@@@@@@@@@@@@@@\`                     *=@@/\@OO@@@O^,^   ]//        =@@@
//                       =O =@@@@@@@@@@@^=^     ,oOO@@@@@@@@@@@@@@@@/[[[[ ,[[[`           *o               O^  ]OOOo[`            =@@^
//                       =O^ @@@@@@@@@@@^=^    =O/[[   ,@OOO@@@OO ,      ,`               =O                                      =@@
//                        O  =@@@@@@@@@@^         * *    OOo\oooOOOO/[[`                  =O*                                     =@^
//                        O   O@@@@@@@@@^                                                  O^                                     /@\
//                             @@@@@@@@@^                                                  ,O                                     O/**o`
//                        ,/[      O@@@@^                                                                                         ooO  O^
//                        O]/oOOO\`,@@@@/                                                                                         OO/^ =O
//                       =/  /oo^  *O@@@\                                                     *                                  =O` [* \^
//                       O^ *=/*    =/O@@`                                                   O^                                  =@@    *`
//                       O^ **    =@@O =@\                                              *   =O                                   OoO\    `
//                       O^  *   OO`/O  O@`                                            [    O^                                   O^*O^  *
//                       O\    ,O/*]oO*,*O\                                                 o                                   =O`*O^ *^
//                        O^   O^***/o^O^ O^                                                [                                   /o,O^  /
//                         O^  \O]***o*O^  \                                                                                    OOO`  /`
//                          O\   [OO]*=O`                                                                                      =O^   /`
//                           \\     ,OO]\O\                                                                                    /@\` /^
//                            ,O`      ,OO/`                                                                                   O@@^/^
//                              ,^        **                                                                                  =^  /^
//                                \`                                                                                          OOO/`
//                                 ,\`                                                                                       ,^
//                                   ,o]    ]OO\                                            *]]]]]]*                         /`
//                                      [OOO`  ,O`                              ]]oo[                                      ,O
//                                               O\                      \OoO/[                                          ,O^
//                                                OO`                                                                  ,O/
//                                                =OO\                                                                //
//                                                 \OoO`                                                            ,O^
//                                                  O\,/*                                                          /OO`
//                                                  =O^*=\`                                                      /O/oO
//                                                   Oo**,[OO]                                                 ,O/**\O
//                                                   \o,***/^,\OO]                                            O/*** =O
//                                                   =O^****o\***[\OO]                                      /O`***  =O
//                                                    O^****,o\******[\OO]                                /O`****   =O
//                                                    O^*****,oo*****,]^,OO@O]                          =O/*****    =^
//                                                    =o*******oo*,***=^**^*\O@@O`                    ,@O******     =^
//                                                    =O********oo******/**oo\]\O@@@\`              ,O@O******      =^
//                                                    =O ********OO********`/oo*o/\O@@@@O]]` ]]]/O@@@@/******       =^
//                                                    =O  ********OO********,]ooo/oo\oOO@@@@@@@@@@@@O`*******       =^
//                                               ]OO@@@/    *******\O*********,[oo/[=oo\,\O@@@@@@O/`*,`****         =^
//                                         ]O@@@@@@@@@/O\     ******\O`*********,][[`,/\*o`**********o/***          =O
//                                       /@@@@@@@@@@@@^  ,[,o]`******\O\****************************/o`**           =^
//                                      =@@@@@@@@@@@@O          [\OOO\OO\**************************=o^**            ,O
//                                     =@@@@@@@@@@@@@^              *[[O@@OO]]*********************oo**            ,]oO\
//                                    =@@@@@@@@@@@@@@\                 ****,[O@@OO]***************=o^*          ]O/`  @@@@\
//                                   =@@@@@@@@@@@@@@@@@\                 ********[\O@OO]`*********,^*     *]/O/[      @@@@@@@O`
//                                  ,@@@@@@@@@@@@@@@@@@@@\                  **********,[OOO\`******   ]o[`            O@@@@@@@@@@\`
//                                 =@@@@@@@@@@@@@@@@@@@@@@@\`                  ************,\oo]`*,O/                 \@@@@@@@@@@@@@@\`
//                                =@@@@@@@@@@@@@@@@@@@@@@@@@@@`                  ************** [O/ ]/O]`             =@@@@@@@@@@@@@@@@@O`
//                               /@@@@@@@@@@@@@@@@@@@@@@@@@@@@O                     *******,O@@@@@@@@@@@@@@`           @@@@@@@@@@@@@@@@@@@@O`
//                             ,@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@^                        *=O@@@@@@@@@@@@@@@@@O          =@@@@@@@@@@@@@@@@@@@@@@\
//                            =@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\                        /@@@@@@@@@@@@@@@@@@@@@`        =@@@@@@@@@@@@@@@@@@@@@@@@\
//                           O@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@^                      O@@@@@@@OOOOOOOOOOOOOO@@^        O@@@@@@@@@@@@@@@@@@@@@@@@@\
//                         /@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@O                    ,@@@@@Oooo***************\@\      ,O@@@@@@@@@@@@@@@@@@@@@@@@@@@\
//                        ,@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@O      ,]`         ,@@@@Oooo\/ooOOO@@@@OOOO],O@@\    OO*O@@@@@@@@@@@@@@@@@@@@@@@@@@@@`
//                        =@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\     / ,OO]     ,@@@@@@OOOO@@@@@@@@@@@@@@@@@@@@\   \O\=@@@@@@@@@@@@@@@@@@@@@@@@@@@@@O
//                      /O@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@^   /  ///O    ,@@@@@@@@@@@@@@@@@@@@@@@@@@@OO@@@\  ,OO^O@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\
//                  ]O@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@O =/ =O=O    ,O@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOO@@\  =OO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\
//             ,]O@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@O[OOO^O^    O@@@@@OOO@@@@@@@@@@@O/[O@@@OOOOOOOO@^  [`=@@@@@@@@@@@@@@@/[\@@@@@@@@@@@@@@@@]
//        ,]@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\   [[    /@@@@OOOOOO@@@@@@@@@@@O^*O@@OOOOOOOOO@^   O@@@@@@@@@@@@@@@@@`  ,\@@@@@@@@@@@@@@`
//  ,]O@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\       =@@OOOOOOOOOOOO@@@@@@@@Oo^o@@@OOOOOOOOOO^  OO@@@@@@@@@@@@@@@@@@`    \@@@@@@@@@@@@@`
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@^     ,@@OOOOOOOOOOOOOO@@@@@@@@O\oO@@@OOOOOOOOOO =@OO@@@@@@@@@@@@@@@@@@@`     \@@@@@@@@@@@O`
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@`    O@@OOOOOOOOOOOOOO@@@@@@@@@OO@@@@@OOOOOOOOO\/@OOO@@@@@@@@@@@@@@@@@@@@\      \@@@@@@@@@@@]
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\`=@OO@@OOOOOOOOOOOOO@@@@@@O*\@@@@@@@OOOOOOOOOOOOOOO@@@@@@@@@@@@@@@@@@@@@\       \@@@@@@@@@@\
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OO@@OOO@@OOOOOOOOOOO@@@@@@@Oo*,O@@@@@@OOOOOOOO@OOOOOO@@@@@@@@@@@@@@@@@@@@@@@`      ,\@@@@@@@@@O
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOO@OOOOOOOOOO@@@@@@@@O\*\@@@@@@@OOOOOOOOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@\       ,\@@@@@@@
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOO@OOOOOOOO@@@@@@@@@O^/O@@@@@@OOOOOOOOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@\        [@@@@
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOO@@OOOOOO@@@@@@@@@O^*oO@@@@@@OOOOOOOOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@O`        ,\
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOO@@OOOO@@@@@@@@@@^*=O@@@@@@OOOOOOOOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\`
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOO@@OOOO@@@@@@@@@o**O@@@@@@@OOOOOOOOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@O]
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOOOO@@O@@@@@@@@@@O**oO@@@@@@OOOOOOOOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@O\
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOOOOO@@@@@@@@@@@O**=O@@@@@@@OOOOOOOOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOOOOO@@@@@@@@@@O^*=O@@@@@@@OOOOOOOOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOOOOOO@@@@@@@@@O^**o@@@@@@@OOOOOOOOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOOOOO@@@@@@@@@@o**oO@@@@@@@OOOOOOOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOOOO@@@@@@@@@@O`*=O@@@@@@@OOOOOOOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOO@@@@@@@@@@@O^*=O@@@@@@@OOOOOOOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOO@@@@@@@@@@@@O^*,O@@@@@@@@OOOOOOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOO@@@@@@@@@@@O^**o@@@@@@@@OOOOOOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOO@@@@@@@@@@@@^**\O@@@@@@@OOOOOOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOO@@@@@@@@@@@@^**=O@@@@@@@@OOOOOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOO@@@@@@@@@@@@^**=O@@@@@@@@OOOOOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOO@@@@@@@@@@@@\***o@@@@@@@@@OOOOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOO@@@@@@@@@@@@o***oO@@@@@@@@OOOOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOO@@@@@@@@@@@@@O***=O@@@@@@@@OOOOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOO@@@@@@@@@@@@@O***=O@@@@@@@@@OOOO@OOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOO@@@@@@@@@@@@@O****O@@@@@@@@@OOOOO@@OO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OO@@@@@@@@@@@@@O^**=o@@@@@@@@@@OOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@O@@@@@@@@@@@@@O****=O@@@@@@@@@OOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@O^***=O@@@@@@@@@@OOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@O^***=O@@@@@@@@@@OOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@

type VulgarTycoon struct {
	heroInfo
}

func (f *VulgarTycoon) Init() {
	f.Id = 18
	f.Rarity = NRarity
	f.MoneyLimit = 110
	f.CurrentMoney = f.MoneyLimit
	f.Name = "土豪"
	f.Motto = "我穷得只剩钱了，我捐3金币"
	f.SkillInfo = []SkillInfo{{
		Name: "一掷三金",
		Desc: "捐赠金币数量至少为3"},
	}
}