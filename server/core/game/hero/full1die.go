package hero

//
//
//                                                                                .]/
//                                                                             ./\@o\    .//\.
//                                                                            /`@OooO/[[`/`..,^
//                                                                 .\       ,/,@oo@/....@`.....\`                    .
//                                         .\@\.                   =`\\` .]@.,@oO/...  @`        ,\].            ./[=`
//                                           .\.\`   `            ,^.  ...=^ /O@^.... =^          ...\/[@]`     /`..O
//                                            .\..@/[.\]         ,^.. ...,/..@o*..... @.             .,\  .,\ ,/   .\.                  .]/O[[@`
//                                             =^  \^. ..\`    ./\.     ,@  =@` ..... @.            ...,@.. .\^.     \.             ,/[.  ..=^
//                                             .^...\`..=@o@]]@o@`     .[   =^........[              .`.,@....=^.   ..=\          //...    /.
//                                             .^ ...@  .\Ooooo@^.          . .      .              ..\`.,^ ...=^..   ..,\`     /`@`.   ..@
//                                             =`.. .\^  =@oooo/.                                     .@..                ..[[\/,@.     .=`
//                                            ,/.    =^..=@ooo@`.                                      ,\.                .,^,/ @..      O
//                                           ,/..    =^..=@ooo@..                                     ..=^                ..@=`=^.      .^
//                                         ,O`      .@.  =@ooO^.                                        .@. =^...         ..@...@..     .\.
//                                      ./[.        ,^...=@oo@^         ../.                            .=^..\^.          ..=^           .\
//                                     ,^=/..           .=@ooO^        .,@...                            .@`..,\.          .@.  ..       ..\`
//                                    ./=^               .@ooo.       .=@`...                             ,\...,@..       ..@.   @`       ..=^
//                                   ,/=^         ... .  .\Ooo.     ..=@^                                 .@... .@.       ./^   ..@`.       .=`
//                                  ,^ @...        =^ . ..,@oo.     .=o@.                                 .@..   =^.      =^... ..\@`.        O
//                                 =`.=^          .,@..    =Oo.     .o@^                                  .@.   .=@.            ..=^@`.    .. =`
//                                .^  =^...       . \`    . =O*   ..=o@^.                               ../^     =@`.             .@,\    ..*./[^
//                                =^  .@...       ...\`.     .\]....=o@^                                .=^..   .@o^              ./.@.   ..=`..=^
//                                =^  .\^..         . ,@`..    .,@\.=oo@                                ,@  .  .@Oo*            ..=^ @...   =^  .O
//                                 O.  .@..            .,\]]...  ].\@ooO\.. ..,/@`.......  ....        .@..  ./@oo/.            .....@...   .O^ .=.
//                                 O.   =^.        .]@/[.....\\`. .[@@oo@^..//. .,//\//[...,[\@]... ...@` ..,@oooo..   ,@\]]..     .=^.    .]O^..=.
//                                /`..   =^      ,@`.   .`.... ,\.. .,@oO^,/...//.... ...     ..,@` ,@/^. ,@ooO@/....]`    ..,\\. .=^..    .@O^..=.
//                               O.=O^.       .,@`        ,\`.  .\`.. .@@@^`.,/..........         .`. =\.=@oO@`..]O@`..     ....,\..       =@O^. @
//                               =^,oO\.     .=/.....]]]]/].,@  ..\`   .@@/..@.                    .@` ,@Oo@^.]oO@` ...                @. .@o@^ ,^
//                               .^=@oO@.   .=@.../oO@@@@@@@\...........=/.....       ,/@@\]         \^.=@@^/oo@^......             ..=^../Oo@../. ,
//                                O.\Oo^@`. .@^.*oooooooO@@@@@.....   ..@`......,.../@@@@@@@@\.     ..\^.@@/oo@^.     . ,@@@OOO`.  .\@`..=@o@`.=^ ,@
//                               .^ .@oo.\^ =/@ /oooooo@@@@@@@\..      .@..   .@...@@@@@@@@@@@@`      .@.=@ooO/..    .=@@@@@ooooo*...@  ./o@`.,@/OO^
//                               =^...\@\...=^\`\ooooo@@@@@@@@@.      ..@.  ..@.. =@@@@@@@@@@@@@^.    .@.=@ooO^.    .=@@@@@@@oooo@`..\.,oO/.,/@Oo\/
//                                =@\...\\..,\.@,ooooooo@@@@@@@..     ..@`   =^. .o@@@@@@@@@@@oo@   .=O\.O@ooO\..    =@@@@@@OOooo@^..=Oo@\ooo@OoO/
//                                 =`,@`..\\.@`.\@ooooooO@@@@@^..     ./\\.`.=^..*oO@o@@@@@@oooo@..  =Ooo@@ooo@`.    ,@@@@@OooooO/...@ooooooooo@`
//                                 .^ ..\\...,@. .\@ooooooo@@`...   ../`/@^\^.\`..ooooooooooooo@^.. .@ooo@@Ooo^..   . ,@Ooooooo@^...,@ooooooooo@.
//                                 .^    ...\\/@. ...[..*[[.. .,@` .,@./@@@`=^....,\Oooooooooo@`.   =Ooo@@@@Ooo`,\.   ...[/[[[..   .@Ooooooooo@@].
//                                 =\..      ..\@`...     ..,//...,/`.@@@@@@\,@`.  ..,[@@@@/`..   ../oo@@@@@@@oo`.\\`...    .../^.,@@ooo@Ooo@@@`
//                             .`  =.\\..  ,@`  .\@]   =@[[`.. .]@`,/@@@@@@@@@`,@`..              .oo@@@@@@@@@@Oo\...[@@]]]]@/`../@@oO@oooO@O/
//                              @[[`. .,`    .[@].,@@@\`.. ..,[.]/@@@@@@@O@@@@@@\.\`.         ..,/o@@@@@@@@@@@@@@@O\`..     ..]@@@O@@ooooooo/
//                              .\.       ,\`.....`/[[@@@@@@@@@@@@@@@@oooOOO@@@@@@@@]`.......,/O@@@@@@@@@oO@@@@@@@@@@@@@@@@@@@@@@ooooO@Ooooo\  ,`
//                               .\..     ,@=\.=^.@.. ...[@@@@@@@Oooooooo@O@O@@@@@@@@@@@@@@@@@@@@@@@@@@OoooooO@@@@@@@@@@@@@@@@@@ooo@@o@@@ooooO/
//                                 \`..    .\\\\=@^ .     .../o@oooooooooo@o@o@oO@@@@@@@@@@@@@@@@@@@@oooooooooooooO@@@@@@@Ooooo@@@O@O@O@ooooO^
//                                  ,@O@\]`...[O,@@`      ../oO^...[ooooooo@O@o@ooO@O@@@@@@@@@@@@oooooooooooooooooooo@@oooooooO@@@O@O@Ooooo@O^
//                                   ,OoooO\..../@\,^..  ..@@@@@`.......[\ooOOOO@O@@oooooooooo@@ooooooooooooooooooooO@@@OooooOOOo@@@@O@oo@@O/
//                                   .@@/[@O@`=@ooO\.^.,//`.. ..,@`.       .,o@O@O@@oooooooooO@@@@O@Oooooooooooooo@@ooooO@Ooo@o@O@@o@OoooO@@.
//                                  ,@... ,@/@@ooooo^//...]]@].....\\`...    ..[\o@@oooooo@oO@@@@@@oOOOoooooooo@@ooooO@@ooo@Oo@o@@@@oo@@@@@@
//                                  @`.....=@.=^.,oO@*/@[..... [\\`...\@`.... ...,@@@ooooo@@@@@@@@@ooooooooO@@oooO@/..oooO@@O@Oooo@@o@Ooooo@.
//                                 =^ .@\...\^@^ ./@@\...   ..... ,\\]...,\\].   .=@@@@ooooO@@@@@@OooooO@@oo@@O\. . .=ooooooO@@o@o@@@OoooO@o@
//                                 =^...@@...@@^ ,[@OO\]/@/[[@@\]]...../oO@@],[@]..@@@O/ooooo@@@@@@@@@O@@@Ooooo^]]/@@@OO@@@@@@@ooo@@Oooo@@oo@^
//                                 @.  =@o@^.=@^.   *oo^O]]@@@@@@@@@@@@@@@@@@@O@@@@@OoooooooooO@@@@@@@@@@@@@@@@@@@@@@@@@@@ooOOoooo@@ooo@@oooO^
//                                .@   /@oo@\.@^      .,[[[[..=^.\O@OO@@@.,@@O@@@@@@ooooooooooo@@@@@@@@`.=@@O@@@@.=@..,[[\\ooooooo@OoO@O@Oooo@
//                                .@  .@OooO@@,/.      .=@\.. =@@O@@@OO@/....[@@@@[..\oooooooo=@@@@@@.. .=@@O@@@OO@/.  /@@oooooooO@o@@@O@Oooo@
//                                .@. .@Ooo@^,@=@      ,[,@@\..\@@OO/@@/..]@@@O@@^...=ooooooo`=@@@@@@@@\`.=@@OOO@@/../@@OOOoooooO@o@@@@O@Oooo@
//                                .@. .@Ooo@` ..@`     ....@@@\` ,[[[\]@@Ooo@@@@^....=ooooooo..=@@@@@@@oO@@\][[[..,/@@@Ooooooooo@Oooo@oO@oooO^
//                                 =\.  =@O@\...=^.   =^... ..[@@@@@@@oo[@@@@@@^.    =ooooooo...,@@@@@@@Oo/oO@@@@@@[..=Ooooooooo@oooo@O@Oooo@`
//                                  ,@` ..,\@@@`=@... =^......,...........\oO@^..   .=ooooooo..  .\@@@@@@oooooo.  ...  ,@OOooooO@ooo@@ooooo@^
//                                    \\.   ..[@.@^.  =^                  .=O^       /oooOooo.@....,@@@@Oooooooo.     . =Oooooo@Oo@@oooooo@^
//                                     .\\.      =\                       .*@.... =^.@O@@@O@@^=^   .=@@@ooooooo[        =@ooooo@O@Ooooooo@`
//                                       @@\../\..@..                      ,`.  . =^@oooooo@O^=^    =@@oooooo[.         ,Ooooo@@ooooooo@/
//                                       =/O@\[[..@^.                           ..@,@ooooooOo.,@. ../@oooo/.          ..@ooooo@Oo@[@oO@`
//                                        \`\o@\..@\.                           ..[.[oooooooo..@...=@ooo`..            =@ooooO@ooooo@`
//                                         \`.\oOOo@. ..@..                         *oooooooo. ....Oo`.               ./Ooooo@@ooO@`
//                                         ,@\..ooo@^...=^.                         .\ooooooo. ..,/`.                 .@ooooo@,[.
//                                      .//.,\@\.,o@^..  @.                          .,oooo/..../..                  .=@ooooO^
//                                     //../@[... ,o@..  =^..                   ..=\]^...[..,`/@..                  ..@ooooo@`
//                                   ,@..//.]@[`.  =@^.  .@.                       ,@\` ..../@/.                    .=@ooooo/
//                                  ,/.//.//...    .o@. ..=^.                       .,@@@]@@@`.                      /Ooooo@`
//                                 .@.@^.@`....   ..=@^.   @.                         .,\@@`                        .@ooooO/
//                                 =^@^.. .....     .@@^. ...                   ../@@@@`..,/@@@`                   .=OoooO@.
//                                 =@^..             @O@^..                 ..,/@ooooooooOooOooo@@]`..            ..\oooO@@.
//                                 .@^.           .,@@Oo@\.               ..=Oooooo@@@@OooO@@@Ooooooo.            ..=oo@@@@\`
//                                  =\.          /@`.=@ooO@`.                =@@/[... ..,[.....,\@OO^..            .o@@@@@@.,\`
//                                   =\.      ..@^.  =@oooo@\..               ...          .       ..             .O@@@@@@@. .,@.
//                                    =\      .@` ..,@@oooO@O@\.            ..,@`               ./@..            /@@@@@@@@@\  ..@
//                                     =^..   =@. ..@@@\ooooo@@@`.            ..\@@@@@@@@@\]]]]@@@.           ./@@@@@@@@@@@@^ ..=^
//                                   ./O@`.   .@` ..@@@.\ooo@@Oo@@`.            .,\@@@@@@@@@@@@@`.          ./@@@@@@@@@@@@@@. ..@.
//                                 ./`=o@\.   .@@`  .@@^.=ooooo@@@@@`.            . ,oooo@@..              /@@@@@@@@@@@@@@@`. ./@.
//                           ..]]@@/. .=O@    =@,@^...\@`..\oo@@OoO@@@`..         ..oooO@@O..            /@@@@@@@@@o@@@@@/   ,@`=^
//                    ,//`..... //... ..=@.   =@./@@.. .\\`..\@Oo@@o@@@@`.        ..=oooo.            ./@@@@@@@@@Oo@@@@/. ../@@`=^
//                ,]@[....]]ooo@^....   .]@/ .=\@ooo@\....\@`..[@@oooO@@@@`..        .\/.         .../@@@@@@@@@@@OO@@/....,@Ooo@@^
//           ,]@/`...]]]@@ooooo@^     ./@`....=^@ooooO@`    ,@\ .,\@OooO@@@@\.        ...        .,@@@@@@@@@@@O@oO@/   .,@@Oooo@O^
//         .@......   .[@@@@OOo@^.   ,@`    ..@^\Oooooo@@`. ..,@\...,ooooo@@@@\..              ./@@@@@@@@@@@o@@@@`  ..,@@OOOooo@,@
//          @` ,\@\`.   ...[@@@@@.. //..    . @^=@ooooooo@@`.....\@]...\oooo@@@@@\`..      .,/@@@@@@@@@@@@OoO@/.....,@@OOOOoooO^.@
//           @`....,\@]..   . ,\@\          .,@@@@OooooooO@@@`.....,\@`..[oooo@@@@@@@@@@@@@@@@@@@@@@@@@@Oo@@`.. ..,@@@OOOOoooo@O\@
//           .@..   . ..    .....\\.    ..]@/`@@ooO@oooo@OoooO@`..... [@]..[oooo@@@@@@@@@@@@@@@@@@@@@@@O@/... ..,@@OOOO@Oooo@@OoO@
//            ,@..               .=@` .]@/... @@oooo@ooo@OooooOO@\......,@\`.\ooooOO@@@@@@@@@@@@@@@@@@@`.    .]@@@OOOOO@Ooo@OOOoO@
//             =^..      ..,]]/@@@@^..@^...   @@ooooooooO@oooooo@@@@`..   .\@`,ooooooo@@@@@@@@@@@@@@@`    ../@@@OOOOOO@@OOOOOOOo@@
//              @.. ./@@/[`.,]]]]]@\]/@^...   @@Ooooooooo@ooooooooO@@@]..   .\@/ooooooO@@@@@@@@@@@/....  ,@@@@OOOOOOOO@@OOOOOOOo@@^
//              ,\. ,]@@[[.    ,@`...@@^.     @O@oooooooO@@Oooooooooo@@@\.   ..\@oooooo@@@@@@@@@@`    .]@@@OOOOOOOOOO@@@OOOOOOoO^@@\
//               ,[[.        ,@`.   .@@@... ..@@O@ooooo@oooo@@ooooooooo@@@\.....@Oooooo@@@@@@@@@.   .,@@@OOOOOOOOO@@OOOO@OOOOO@@@@@@@]
//                        ,//.      =@@@..  ..@Ooo@Oooo@ooooooO@oooooooooo@\    =@ooooo@@@@@@@@^.  .=@OOOOOOOOOO@@OOOOOO@OOOO@ooo@@@.@@@].
//                    .]@[..        =@@@^.    \@oooO@oo@Ooooooooooooooooooo@`   .@ooo@O@@@@@@@@`. ..@OOOOOOOOOOOOOOOOOOO@OO@@OOoO@@@..\@oO@\`
//              .]]@/[. .           =@@@@.  ..=@oooooooO@oooooooooooooooooo@\.  .\@ooO@@@@@@@@@.. .=@OOOOOOOOOOOOOOOOOOO@OOOOOoo@@@@...,@OoooO@\`
//      .]]/@@[[....                =@@@@^.  .,@oooooO@OO@OooooooOo@Oooooooo@.. .=@oo@O@@@@@o@^.. .@@OOOOOO@@OOOOOOOOO@@O@@OOOoo@@@^. ...\@ooooooO@@\`
//  .[@].                           =@@@@@    .=@oooo@ooooo@Ooo@OooooO@Ooo@o@^  .,@ooo@@@@@oo@^   ,@O@OOO@OOOOOO@OOO@OOOOO@OOOo@@@/,\@]....\@oooooooooo@@]
//      ,@\.                        .@@@@@`.]//`=@ooo@oooooooO@oooooooo@Oo@@@\  ..@Ooo@@@Oooo@`   =@O@OO@OOOOOOOO@@OOOOOOO@OOo@@@@.    .,[@\][[\ooooooooooO@\`
//         ,@]. .                   .@@@@@`..   .,@oo@Ooooooo@oooooooooo@oO@O@. ..@Oooo@@@ooo@. ../@@@O@OOOOOOOOOO@OOOOOOO@OO@@@@`         ...,[\@@OoooooooooO@.
//            [@`..                 .=@@/..       ,@OO@oooooo@oooooooooo@Oo@o@.  .=@ooooooooO@  . @O@OO@OOOOOOOOOO@@OOOOO@@O@`=@^.                  . ...,\@/.
//              ,\\..           .,/. ,@`.         . \@@oooooo@oooooooooo@oo@o@^  .=@oooooooo@/   .@O@OO@OOOOOOOOOO@@OOOOO@@@..=@.                   ... ,@`
//                 \\....   . ,@/..                 .,@Oooooo@ooooO@Oo@@ooo@O@^.. =@oooooooo@^  .=@O@OOO@@OO@@OOOO@OOOOOO@^ .                       ..]@`
//                   ,@`..,/@[.                       .\@ooooO@oooooooooooo@O@^   =@oooooooo@^   =@O@OOOOOOOOOOOO@@OOOO@@..                       ..,@.
//                     ,@^...                         . ,@@oo\@Oooooooooooo@O@^.. =@oooooooo@^   =@O@OOOOOOOOOOOO@OOO@@`...                       ,@`
//                       ,@`.                           ..,@Ooo@@oooooooooo@o@^   .@oooooooo@^   =@O@OOOOOOOOOO@@OOO@`.....                     ./^
//                         =\..                 ...,]/@@@@@@@@@@@@@@@Oooooooo@^   *@oooooooo@^   =@OOOOOOOO@@@@@@@@@@@@@@@@\].                 ,/.
//                           \\ .           .,]@@@OooooO@oo@oooooOOoooO@@@@Oo@^   =@oooooooo@^  .,@OO@@@@OooooooooO@oooooooooO@@\`.          ./^
//                            ,@`..     .,/@@oOOoooooooo@@@@ooooo@@o@@ooooooO@@]..=@oooooooo@^..]@@OoooO@oo@@ooooo@@O@@ooooooooooO@@`..     ,@`
//                              =\    ./@Ooo@@o@@ooooooooO@@@@@@OOOOooooooo@@oOO@@@@ooooooooO@@@@Ooo@oooO@@OooOO@@@@@@OOooooooooOOooO@\.  .,@.
//                               ,@`,@@oooooo@@@ooO@@[`.           .[\@@Ooo@@O@@ooO@[@oooo@@@@oO@OO@@ooO@@/[.            ,[@@Ooo@Ooo@oo@@`=/
//                                 \@ooooooooo@@[.                       .[@@Oo/ooo@. =@@/  @OooooO@@[`                       ,\@O@@ooooo@/
//                                  \@ooooO@/.                                ,\@/`    ..    ,\@/[.                              ,@OooooO/
//                                   =@o@/.                                                                                        ,@ooo@
//                                    \/.                                                                                           .@o@`
//                                                                                                                                   .@/
//                                                                                                                                    ,`

type Full1Die struct {
	heroInfo
}

func (f *Full1Die) Init() {
	f.Id = 13
	f.Rarity = RRarity
	f.MoneyLimit = 110
	f.CurrentMoney = f.MoneyLimit
	f.Name = "富一代"
	f.Motto = "可惜父亲见不到我现在的样子了"
	f.SkillInfo = []SkillInfo{{
		Name: "白手起家",
		Desc: "使用富一代取得游戏胜利后\n该角色升级为富零代"},
	}
}
