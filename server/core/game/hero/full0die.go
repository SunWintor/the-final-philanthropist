package hero

// https://mzh.moegirl.org.cn/%E8%B4%AB%E7%A9%B7%E9%85%B1
//
//
//
//                                                                                              ..]]]@@@@@@@@@@@@@@@@@@@@@@@\]]]..
//                                                                                      .]]@@OOOOOOOooooooOOOOOOOOOOOOOOOOOOOOOOOOO@@@@]].
//                                                                                .]/@[`,oooooooooooooooooOOOOOOOOOOOOoooooooooooOOOOOOOO],[@]`
//                                                                            ,/@[..]ooooooooooooooooooooooooooOOOOOOOOOOoooooooooooooooooOOO].,\@].
//                                                                        ,//`..]ooooooooooooooooooooooooooooooooooOOOOOOOOooooooooooooooooooooo`.[/\\`
//                                                                    .//`..]oOOooooOOOOOOOOOOOOOOooooooooooooooooooooOOOOOoooooooooooooooooooooooo\.,\/@`
//                                                                 .//.../OOOOOOOOOOOOOOOooooooooOOOOOOooooooooooooooooooOOOOooooooooooooooooooooooooO\`,],@`
//                                                              .//`./OOooOOOOOOOOOOOOOOOoooooooooooooOOOOOooooooooooooooo^,OOooooooooooooooooooooooooOOOO`..\\.
//                                                            ,@`.]OOoOOOOO@OOOOOOOOOOOOOOOOOOooooooooooooOOOooooooooooooooo`,OOoooooooooooooooooooooooOOOOOO`.,@`
//                                                         .//../oOOoOOOOOOOooooOOOOOOOOooooooOOOoooooooooooOOOooooooooooooooo`,OoooooooooooooooooooooooOOOOOOOo`.\]
//                                                       .//.,ooOOooooooOOooooooOOooooOOOoooooooOOOOOooooooooooOOoooooooooooooo\.\`ooooooooooooooooooooooOOOOOOOOO\.\\
//                                                     .//./OOOoooooooooOOoooooOOooooooOOoooooooooOOOOOOoooooooooOOoooooooooooooo\.*OoooooooooooooooooooooOO@OOOOOOO\.\^
//                                                    ,/.,OOOO\ooooooooOOooooooOooooooooOOooooooooooO\/Ooooooooooo\=ooooooooooooooo\.=oooooooooooooooooooooOO@OOOOOOOO^.\`
//                                                  .@`,oOOOO\OoooooooOOooooooOOooooooooOOooooooooooooo\\ooooooooooo=oooooooooooooooo`,OooooooooooooooooooooOO@OOOOOOOOO`,\.
//                                                 =/./OOOOOOooooooooOO*ooooooOOooooooooOOOoooooooooooooo^\ooooooooooOOooooooooooooooo\.\oooooooooooooooooooOO@@OOOOOOOOO\.\`
//                                               ./`,OOOOOOOooooooooOO`/ooooooOOoooooooooOOOooooooooooooooo=oooooooooo\Ooooooooooooooooo^,OooooooooooooooooooO@@@OOOOOOOOO\.,\
//                                              ,/./OOOOOOOooooooooOO^,ooooooOOOoooooooooOOOoooooooooooooooo/ooooooooo\,Ooooooooooooooo/.*.=oooooooooooooooooOO@@OOOOOOOOOO\*.@.
//                                             ,/./OOOOOOoooooooooOOO.oooooooOOoooooooooooOOOoooooooooooooooo/ooooooooo\,Ooooooooooooooo..,.=oooooooooooooooooOO@@OOOOOOO[\o^..@`
//                                            =^.=OOOOOO*ooooooooOOO`=oooooooOOoooooooooooOOOooooooooooooooooo\=oooooooo^OOooooooooooooo...`.\ooooooooooooooooOO@@OOOOOOO*`\o*..\`
//                                           =^..OOOOOO`oooooooooOO/*ooooooooOOoooooooooooOOOOoooooooooooooooooo,oooooooo`\Ooooooooooooo...*..oooooooooooooooooOO@@OOooOO`..,^...\`
//                                          =^..,OOOOO`/ooooooooOOO`=ooooooooOOooooooooooooOOOooooooooooooooooooo.\ooooooO.=Oooooooooooo..^.^.=ooooooooooooooooOOO@OOooOO^.\..^*..@.
//                                         =^.*.,..OO^=oooooooooOOo*oooooooOOOO\oooooooooooOOOOoooooooooooooooooo..ooooooo^.OOoooo/,oooo..o.,..oooooooooooooooooOO@@OoooO`=Oo..=...@.
//                                        =^.=.*^.,\@^oooooooooOOOoooooooooOOOO=ooooooooooooOO/ooooooooooooooooooo..oooooo^.=OOoooo*,`.,..=^...\oooooooooooooooooOO@OOoo..ooo\..^..,\
//                                       ,/.,O.*...=O=oo[/[\oooOO/`,[[[[[[[\[[^...[[[[[[oooooOO/oooooooooooo/..,[`...o`......OOooooo`.=`..=o...,ooooooo[[.,[\ooo/[\/OOoooooooo^.\`..=^
//                                      .@..O^..*.]@``/o\.....,/Oo*]]]]]/ooO@O@=oo`.*.**.......`........../oo\..`,o`.=,*,/\^.=OOoooooooo..=o^...oooo\]ooooooooooooOO@Oooooooooo..=*..@.
//                                      /`./@^.,OO@@/oooooooooo@OO`ooooooooO@O@^oooooooooooooOO@`ooooooooooooo`=ooo\..oOOOoO^.OOoooooooo\`/oo\..=oooooooooooooooooOOOOOoooooooo].=\..,\
//                                     =^.=@O^.OOo@@@oooooooooO@@O^ooooooooO@@@\OoooooooooooooOO@^oooooooooooooooooO^.\OOOOOO.=Ooooooooooooooo..,oooooooooooooooooOOOOOoooooooooooo^..\`
//                                    .@..@@O^,OoO@@OOooooooooO@@O,ooooooooo@@@@OoooooooooooooOO@O=oooooooooooooooOOo.=OOOOOO^=@Oooooooooooooo`..ooooooooooooooooooOOOOooooooooooooo^.,@
//                                    /`.O@@OO/OoO@@^OOoooooooO@@@^oooooooooO@@@OOooooooooooooo@@@OOooooooooooooooOOO^=OOOOOO\]@Oooooooooooooo^..=oooooooooooooooooOOOOoooooooooooooo^.\^
//                                   ,/.=@@@OOOOoO@@\OOoooooooO@@@\=ooooooooO@@@@OOooooooooooooO@@@OOoooooooooooooOOO\.@OOOOOOO@@Ooooooooooooo^..=OoooooooooooooooooOOOoooooooooooooo^..@
//                                   /..@@@OOOOOoO@@@/OoooooooO@@@@/Ooooooooo@@@@@OOooooooooooOOO@@@OOOoOoooooooooOOOO.@OOOOOOO@@OOOOooooooooo^..*OoooooooooooooooooOOOooooooooooooooo..=^
//                                  ./.,@@@OOOOOoO@@@@=OOoooooO@@@@OOOOOOOOOOO@@@@@OOOOOOOOOOOOOO@@@@OOOOOOOoooooOOOOO.@@OOOOOO^@OOOOOoooooooo\`.,OOoooooooooooooooooOOOooooooooooooooo`.@
//                                  =^./@@@@OOOooO@@@@@OOOOOOOO@@@@@OOOOOOOOOO@@@@@@OOOOOOOOOOOOOO@@@@@OOOOOOOOOOOOOOO`@@OOOOOO^@OOOOOOoooooooo^.=@OoooooooooooooooooOOOooooooooooooooo^.=^
//                                  =`,@@@@@OooooO@@@@@OOOOOOOOO@@@@OOOOOOOOOO@@@@@@@OOOOOOOOOOOOO@@@@@@OOOOOOOOOOOOOO^@@@OOOOO@@@OOOOOOOOoOOoo^.=@OOoooooooooooooooooOOoooooooooooooooo.,\
//                                  @.=@@@@@OoooOO@@@@@@OOOOOOOO@@@@@OOOOOOOOO@@@@@@@@OOOOOOOOOOOO@@@@@@@OOOOOOOOOOOOO^\@@OOOOO@@@OOOOOOOOOOOOO^.O@@OOoooooooooooooooooOoooooooooooooooo^.@.
//                                  @.@@@@@@OOOOOO@@@@@@@OOOOOOO@@@@@@OOOOOOOOO@@@@@@@@OOOOOOOOOOOO@@@@@@@@OOOOOOOOOOO^=@@@@@@@@@@OOOOOOOOOOOOO^.O@@OOoooooooooooooooooOOooooooooooooooo^.=^
//                                  @.@@@@@@OOOOOOO@@@@@@@OOOOOOO@@@@@@OOOOOOOOO@@@@@@@@@@@@@@@OOOO@@@@@@@@@OOOOOOOOOO\=@@@@@@@@@@OOOOOOOOOOOOO^.@@@OOOoooooooooooooooooOooooooooooooooo\.=^
//                                  @.@@@@@@@OOOOOO@@@@@@@@OOOOOO@@@@@@@OOOOOOOO@@@@@@@@@OOOOOOOOOOO@@@@@@@@@@@OOOOOOOO=@@@@@@@@@@OOOOOOOOOOOOO^.@@@@OOoooooooooooooooooOoooooooooooooooO.,@
//                                  =`@@@@@@@OOOOOOO@@@@@@@@@@@@@@@@@@@@@OOOOOOOO@@@@@@@@@OOOOOOOOOOO@@@@@@@@@@OOO@@OOO=@@@@@@@@@@OOOOOOOOOOOOO.=@@@@OOOOoooooooooooooooOOooooooooooooooO..@
//                                  =^@@@@@@@OOOOOOO@@@@@@@@@OOOOO@@@@@@@@OOOOOOO@@@@@@@@@@@OOOOOOOOO@@@@@@@@@@@OOOOOOOO@@@@@@@@@@OOOOOOOOOOOOo.=@@@@OOOOooooooooooooooooOooooooooooooooO..@
//                                  =^@@@@@@@@OOOO@OO@@@@@@@@@OOOOO@@@@@@@@OOOOOOO@@@@@@@@@@@OOOOOOOOO@@@@@@@@@@@OOOOOO@@@@@@@@@@@@OOOOOOOOOOO^./@@@@@OOOOoooooooooooooooOooooooooooooooO..@
//                                  .@@@@@@@@@OOOOOOO@@@@@@@@@@OOOOO@@@@@@@@@OOOOOO@@@@@@@@@@@OOOOOOOO@@@@@@@@@@@@OOOO/@@@@@@@@@@@@OOOOOOOOOOO`=@@@@@@OOOOOooooooooooooooOOoooooooooooooO..@
//                                   @=@@@@@@@@OOOOOOO@@@@@@@@@@@OOOO@@@@@@@@@OOOOOO@@@@@@@@@@@OOOOOOO@@@@@@@@@@@@@OOO^@@@@@@@@@@@@OOOOOOOOOOO.=@@@@@@OOOOOooooooooooooooOOoooooooooooooO.=^
//                                   =^@@@@@@@@@OOOOOOO@@@@@@@@@@@OOOO@@@@@@@@@@OOOO@@@@@@@@@@@@OOOOOOO@@@@@@@@@@@@@OO\@@@@@@@@@@@OOOOOOOOOOO^.@@@@@O@@OOOOoooooooooooooooOooooooooooooOO.=^
//                                   ,^O@@@@@@@@@OOOOOO@@@@@@@@@@@@@OOO@O@@@@@@@@OO@@@@@@@@@/@@@@@@OOOO@@@@@@@@@@@@@@OO@@@@@@@@@@@OOOOOOOOOOO^.@@@@OO@@OOOOoooooooooooooooOOoooooooooooOO.=^
//                                    @=@@@@@@@@@@OOOOOO@@@@@@@@@@@@@OO@@*,@@@[@/\[[,@@@@@@@@@@@@\]*`[@@@@@@@@@O@@@@@@@@@@@@@@@@@@OOOOOOOOOOO`=@@@@OOO@OOOOOooooooooooooooOOoooooooooooOO.@.
//                                    =`@@@@@@@@@@@@@@@@@@@@@@@@@\`***[`****.....*****,@@@@@@@@@@@@@@@@@@@@@@@@@o@@@@@@@@@@@@@@@@@OOOOOOOOOOO.O@@@@OOO@@OOOOooooooooooooooOOoooooooooooOO.@
//                                    =^\@@@@@@@@@@@@@@@@@@@@@@@@@@@/................=@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOO^.@@@@@OOO@@OOOOooooooooooooooOOooooooooooOO^=^
//                                    .@]@@@/`\@@@@@@@@@@@@@/ .,@@@@@^..................@@@@@@@@@`. .[\@@@@@@@@@@@@@/@@@@@@@@@@@@@OOOOOOOOOO`=@@@@@OOOO@OOOOoooooooooooooooOooooooooooOO^=^
//                                            .@@@@@@@@@@@@@.....@@OO^..................@@@OOO\@. .  .]@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOO.@@@@@@OOOO@OOOOoooooooooooooooOOoooooooooOO^@.
//                                             =@@@@@@@@@@@@@@\/@@@@O^...................*.....@@@].,@@@@@@@@OOo\**@@@@@@@@@@@@@@OOOOOOOOOOO=@@@@@@OOOOOOOOooooooooooooooooOOoooooooooOO.@
//                                              @@@@@@@@@@`=@@O@@@@@O^.........................\@@@OO/@@@@@@@@OOo^=@@@@@@@@@@@@@@OOOOOOOOOO^=@@@@@@OOOOOOOOoooooooooooooooooOoooooooooOO=^
//                                              ,@@@@@@@\...@,^.@@@`,`.........................,\*\O\.=@@@@@@[,O/*@@@@@@@@@@@@@@@OOOOOOOOOO`@@@@@@@OOOOOOOOoooooooooooooooooOooooooooOO^=^
//                                               \@@@@@@@`..,\,O\*]O.`.....................     =^..[O`....../O`..=@@@@@@@@@@@@@OOOOOOOOOOO=@@@@@@@@OOOOOOOOooooooooooooooooOooooooooOO^/.
//                                               ,@@@@@@@@...=^...,/O^.......................    \...*,[oOOo[`.,/\o@@@@@@@@@@@@@OOOOOOOOOO^O@@@@@@@@OOOOOOOOooooooooooooooooOOoooooooOO^@.
//                                                @@@@@@@@^...,@@@@@o..//`........................\\@@@@@@@OO@@@Ooo@@@@@@@@@@@@@OOOOOOOOOO=@@@@@@@@@OOOOOOOOOoooooooooooooooOOoooooooOO.@
//                                                =@@@@@@@@..]`/[`..,/`..............................,[[[****=*]/OO@@@@@@@@@@@@OOOOOOOOOO^O@@@@@@@@@OOOOOOOOOooooooooooooooooOoooooooOO./
//                                                =@@@@@@@@/O\\^/,.`.................................].,*,/,/,/,/O[@@@@@@@@@@@@OOOOOOOOOO\@@@@@@@@@@OOOOOOOOOOooooooooooooooooooooooOOO,^
//                                                ,@@@@@@@@\//\//\^............................../\\/,/,/,/,/,/,/**@@@@@@@@@@@@OOOOOOOOOOO@@@@@@@@@@OOOOOO@OOOooooooooooooooooooooooOOO=^
//                                                .@@@@@@@@@`=``................................,//,/,/,/,/*/`/`***@@@@@@@@@@@OOOOOOOOOO^@@@@@@@@@@@@OOOOO@OOOOoooooooooooooooooooooOO^=^
//                                                .@@@@@@@@@`.................................../......./.,/.......=@@@@@@@@@@OOOOOOOOOO=@@@@@@@@@@@@OOOOO@OOOOoooooooooooooooooooooOO^=^
//                                                .@@@@@@@@@^ .....................................................=@@@@@@@@@OOOOOOOOOO^@@@@@@@@@@@@@OOOOOO@OOOOooooooooooooooooooooOO^=^
//                                                .@@@@@@@@@@`.....................................................=@@@@@@@@@OOOOOOOOOOO@@@@@@@@@@@@@OOOOOO@@OOOOoooooooooooooooooooOO^=^
//                                                .@@@@@@@@@@@.....................................................=@@@@@@@@@OOOOOOOOOO@@@@@@@@@@@@@@OOOOOO@@OOOOoooooooooooooooooooOOO=^
//                                                =@@@@@@@@@@@^.....................................................@@@@@@@@OOOOOOOOOO@@@@@@@@@@@@@@@@OOOOOO@@OOOOooooooooooooooooooOOO=^
//                                                =@@@@@@@@@@@@\....................................................@@@@@@@OOOOOOOOOOO@@@@@@@@@@@@@@@@OOOOOO@@OOOOooooooooooooooooooOOO=^
//                                                =@@O@@@@@@@@/\\..............=`...................................@@@@@@@OOOOOOOOOO@@@@@@@@@@@@@@@@@OOOOOO@@@OOOOoooooooooooooooooOOOO^
//                                                =@@OO@@@@@@@/ ,@`..............................................,/O@@@@@@@OOOOOOOOOO@@@@@@@@@@@@@@@@@OOOOOOO@@OOOOoooooooooooooooooOOOO^
//                                                =@@OO@@@@@@@/   ,@...........................................]@ooo@@@@@@OOOOOOOOOOO@@@@@@@@@@@@@@@@@@OOOOOO@@@OOOOooooooooooooooooOOOO^
//                                                =@@OOO@@@@@@^     \\......................................,OOoooooO@@@@@OOOOOOOOOO@@@@@@@@@@@@@@@@@@@OOOOOOO@@OOOOooooooooooooooooOOOO^
//                                                =@@OOO@@@@@@^      .@`.................................,/Oooooooooo@@@@OOOOOOOOOOO@@@@@@@@@@@@@@@@@@@OOOOOOOO@@OOOooooooooooooooooOOOO^
//                                                =@@OOOO@@@@@\        =@.............................,/ooooooooooooO@@@@OOOOOOOOOO@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOOOOOoooooooooooooooOOOO^
//                                                =@@OOOO@@@@@@         \@\`.......................]Ooooooooooooo/`*=@@@@OOOOOOOOOO@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOOOOoooooooooooooooOOOO^
//                                                =@@OOOO@@@@@@         @@@@@]................,/Ooooooooooooooo`****@@@@OOOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOOOOOooooooooooooooOOOO^
//                                                @@@OOOOO@@@@@        .@@@@@@@@]........,]@/\oooooooooooooo[******,@@@@OOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOO@OO\oooooooooooooOOOO@
//                                                @@@OOOOO@@@@^        .@@@@@@@@@@@@]]@@@@@`ooooooooooooo/*********=@@@OOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOO@O^oooooooooooooOOOO@
//                                                @@@OOOOOO@@@^         @@@@@@@@@@@@@@@@@@@@`oooooooooo`***********=@@@OOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOOOOO,ooooooooooooOOOO@
//                                                @@@OOOOOO@@@^         =@@@@@@@@@@@@@@@@@@@@=oooooo/`****]]/@@/[[o@@@OOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOOOO\=oooooooooooOOOO@.
//                                               .@@@OOOOOO@@@^         =@@@@@@@@@@@@@/\oooO@.\oO@@@[[[..........*=@@@OOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOOO@O^\oooooooooooOOO@.
//                                               .@@OOOOOOOO@@^         =@@@@@@@@@[..*=oooO@@[.............    ...=@@OOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOOO@O`oooooooooooOOO@.
//                                               .@@OOOOOOOO@@.         .@@@@/[....***ooo@`.......      .     ...*@@OOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOOOO@O,ooooooooooOOO@^
//                                               =@@OOOOOOOOO@           @@^  ...****=ooo@..                 ...*=@@OOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOOO@O^=oooooooooOOO@^
//                                               =@OOOOOOOOOO^           @@@` ..*****=ooo@..               ....**/@OOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOOOO@O^\ooooooooOOOO^
//                                               =@OOOOOOOOO@^           =@@@...*****=ooo@..             .....**=@@OOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOOOO@O`oooooooooOOO@
//                                               =OOOOOOOOOO@             @@@@`..****=ooo@..           ......**=O@OOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOOOO@O*ooooooooOOO@.
//                                               =OOOOOOOOOO^     .,]@@]]/@@/[\`.****=ooo@^.          .....***/o@@OOOOOOO@@@@@O@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOOOO@OO,oooooooOOO@.
//                                               =^OOOOOOOO@.    /@@@@@@@@Oooo`=`****=oooO@.        ......**,o\/@OOOOOOOO@@@@OoO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOOOO@O\,oooooooOO@^
//                                               @^OOOOOOOO/   ,@@@@@@@@@@@@@Oo\\^***/oooo@..     ......**,@/. =@OOOOOOO@@@@@oooo@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOOOOOO\,ooooooOOO^
//                                               @^=OooooO@^ ,@@@@@@@@@@@@@@@@@Oo@`**ooooO@...........]@/`.....@OOOOOOOO@@@@Oooooo@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOOOOOO\,oooooooO@
//                                              .@^=Oooooo@./@@@@@@@@@@@@@@@@@@@@o@`,ooo@@^.**..,]OO[*........=@OOOOOOO@@@@Ooooooo`\@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOOOOOOO\*ooooooO@.
//                                              .@O=Ooooo@@@@@@@@@@@@@@@@@@@@@@@@@@@/oO@@@@@@Oo`*.............@OOOOOOO@@@@@oooooooo/@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOOOOO@OO*\ooooo@^
//                                              .@O*Ooooo@@@@@@@@@@@@@@@@@@@OOO@@@@@@@@OOO@@@@@@@@@@@@]].....,@OOOOOOO@@@@Oooooo\@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOOOO@@OO`\ooooO\
//                                              =@O^ooooO@@@@@@@@@@@@@@@@@@@@@@O@@@@@@@@@@@@@@@@@@@@@@@@@@@]./OOOOOOO@@@@Oooooo@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOOOO@@OO`\oooO@.
//                                              =@O^=ooo@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOO@@@@@oooO@@@@@@@@@O@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOOOO@@OO`=ooO@^
//                                              /@OO*ooo@@O@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOO@@@@ooo@@@@@@@@@OOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOOOOO@@OO\=OOO^
//                                             .@@@O*\oO^@O@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOO@@@@Oo@@@@@@@@@OOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOOOOO@@OO\=OO@.
//                                            .@@@@O^=o@^@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOO@@@@@@@@@@@@@@OOOOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOOOOO@@OO\=O@^
//                                          .//.@@@@O*o@*@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OO@OOOOO@@@@@@@@@@@@OOOOOOOOOOOOOOO@@@@@@@@@@@@@@@@@O@@@@@@@@@@OOOOOOOOOOOOOO@@OOOOO@
//                                        ]@`...@@@@@^O/=@O@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOO@OOOOOO@@@@@@@@@@OOOOOOOOOOOOOO@OOOO@@@@@@@@@@@@@@@@O@@@OO@@@@@OOOOOOOOOOOOOO@@OOOO@.
//                                     ./`......@@@@@O/@@@OOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOO@OOOOO@@@@@@@@@OOOOOOOOOOOOOOO/[\@OOO@@@@@@@@@@@@@@@@OOOOOO@@@@@@OOOOOOOOOOOOOO@OOO@^
//                                    =/.....`..=@@@@@@@@@@@@@@@OooO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOO@OOOOOO@@@@@@@@OOOOOOOOOOOOO/******\@@OO@@@@@@@@@@@@@@@OOOOOOO@@@@@OOOOOOOOOOOOOO@OO@@.
//                                   /`......\**/@@@@@@@@@@OooooooO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOO@@OOOOOO@@@@@@OOOOOOOOOOOo`****.....*,O@OO@@@@@@@@@@@@@@@OOOOOOO@@@@@OOOOOOOOOOOOOOO@@@^
//                                 .@........,@/=@@@@@@@@@Ooooooo@@@@@@@@@@@@@O@@@@@@@@@@@@@@@@@@@@@@OOO@OOOOOO@@@@@OOOOOOOOOO/`*.............,oO@oO@@@@@@@@@@@@@@OOOOOOOO@@@@@OOOOOOOOOOOOOO@@@
//                                ,@.......*,@`*=@@@@@@@@ooooooO@@@@@@@@@@@@@@o@@@@@@@@@@@@@@@OOOOOOOOO@OOOOO@@@@@OOOOOOOOoo`*..................*,@ooO@@@@@@@@@@@@@OOOOOOOO@@@@@OOOOOOOOOOOOOO@@`
//                               /^......**/`***=@@@@@@@@@oooo@@@@@@@@@@@@@@@@oO@@@@@@@@@@@@@@@@@OOOOO@OOOOO@@@@OOOOOOooo[*.......................*@,\o@@@@@@@@@@@@@OOOOOOOOO@@@@@OOOOOOOOOOOOO@\
//                             ,/.......*//.....*@@@@@@@@@@@O@@@@@@@@@@@@@@@@@oo@@@@@@@@@@@@@@@@@@@@@@OOO@@@@@@OOOoooo/`..........................*/@`**\@@@@@@@@@@@@OOOOOOOOO@@@@@OOOOOOOOOOOOO@`
//                           ,/.......*//.......*=@@@@@@@@@@@@@@@@@@@@@@@@@@@@ooO@@@@@@@@@@@@@@@@@@@@OO@@@@[@Ooooooo/*..........................*//......,@@@@@@@@@@@@OOOOOOOOOO@@@@@OOOOOOOOOO@@^
//                         ,/.......,@`..........*,@@@@@@@@@@@@@@@@@@@@@@@@@@@ooo@@@@@@@@@@@@@@@@@@@@@@O[]O@ooo/[**..........................*,@`*.........\@@@@@@@@@@@OOOOOOOOOO@@@@@OOOOOOOOO@@@.
//                      .//......../`*............**,@@@/@@@@@@@@@@@@@@@@@@@@OoooO@@@@@@@@@@@@@@@@@@@/,oo@/`*****..........................*//**.............@@@@@@@@@@@OOOOOOOOOOO@@@@@OOOOOOOO@@^
//                    ./[........//................****@@@@@@@@@@@@@@@@@@@@@@OooooO@@@@@@@@@@@@@@@@/,ooO@`*****.........................*,/`**................,@@@@@@@@@@@OOOOOOOOOO@@@@@OOOOOOOO@@.
//                  ./`......**,/...................**@@@@@@@@@@@@@@@@@@@@@@@@Oooooo@@@@@@@@@@@@@`/o/O@@`****........................*,O`****...................\@@@@@@@@@@OOOOOOOOOOO@@@@@OOOOOOO@^
//                ,O`......***/^....................,@@@@@@@@@@@@@@@@@@@@@@@@@@@@Ooooo@@@@@@@@/`=o//@@@***.........................*/[*****......................,@@@@@@@@@@OOOOOOOOOOO@@@@@OOOOOOO@.
//              .O`.......***@`....................,@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@Oooo\@@@@Oo[,/@@@/..........................*,/*******.........................\@@@@@@@@@OOOOOOOOOOOOO@@@@OOOOO@^
//             .@.......****@.....................=@@@@@@@O@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@Ooo\/@@@@@`.........................**********............................,@@@@@@@@@OOOOOOOOOOOOOO@@@@OOOO@`
//             .^.......***/`....................=@O@@@@@OO@@@@@@@@@@@@@@@@@@@.,\@@@@@@@@@@@@@@@@@/.........................*,*********...............................\@@@@@@@@OOOOOOOOOOOOOOO@@@@OOO\
//              @.....****=^..................../@OO@@@@OO@@@@@@@@@@@@@@@@@@@@*...*\@@@@@@@@@@@@@`........................./`*********..................................@@@@@@@@@OOOOOOOOOOOOOOO@@@OO@^
//              ,\..******@..................../@OO@@@@OO@@@@@@@@@@@@@@@@@@@@@^.......*[@@@@@@/`........................./`**********...................................,@@@@@@@@@OOOOOOOOOOOOOOOO@@@@@.
//               =^******=^.................../@OOO@@@OOO@@@@@@@@@@@@@@@@@@@@@^************............................=/**********......................................=@@@@@@@@@OOOOOOOOOOOOOOO@@@@@@.
//                =^*****@*................../@OOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@***********..........................*,/***********........................................=@@@@@@@@@OOOOOOOOOOOOOOO@@@@@@.
//                 =\..*=/*.................=@OOOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@*********..........................*,/************..........................................@O@@@@@@@@OOOOOOOOOOOOOOO@@@@@@.

type Full0Die struct {
	heroInfo
}

func (f *Full0Die) Init() {
	f.Id = 14
	f.Rarity = SRRarity
	f.MoneyLimit = 5
	f.CurrentMoney = f.MoneyLimit
	f.Name = "富零代"
	f.Motto = "孩子读书的钱一定要攒够"
	f.SkillInfo = []SkillInfo{{
		Name: "传宗",
		Desc: "被淘汰后使用富一代继续游戏"}, {
		Name: "接代",
		Desc: "使用富零代取得游戏胜利后\n该角色升级为负一代"},
	}
}
