package hero

//               ..     ........,O@@@@@@@@@@OOOOOOOOOOOOOOOOO@@@@OOO@@@@@@@@@@@@@@@@@@@OOOO@OOOOOOOO@@@@@@@@@@@/*...
//                     .........*O@@@@@@@@@OOOOOOOOOOOOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@O@O@@@@@@@O@OOO@@@@@@@O`.....
//                     . ....,/O@@@@@@@@@@@@OOO@OOOOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOO@@@@@OOO]`...
//                     ...,OOOOO@@@@@@@@@@@@@@@@@OOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOO/[[.
//                       ....=O@@@@@@@@@@@@@@@@@@O@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OO]]`
//                 .      ..,O@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@O/
//                       ../@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOoOO@@@@O@@@@@@@@@@@@@@@@@@@@@@@@@@@OO/*..
//                      ../@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OO@@@@@@OO***OOOOOOOO@@@O@@@@@@@@@@@@@@@@@@@@@@@O`..
//       .              .=@@OO@@@@@@@@@@@@@@@@@@@@@@@@@@@@OO@@OOOooOOOOOO*...*OooOOoOOOOOOO@@O@@@@@@@@@@@@@@@@OO@@^.
//                     ..OOoO@@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOO`**,OooO^.....[..o`*\OOOOOOOOOOO@@@@@@@@@@@@@@@O`.[.
//                      ...=O@@@@@@@@@@@@@@@@@@@@@@@@@OOO^[[OO/*.....................\Oo`*oOOOOO@@@@@@@@@@@@@@@@\...
//                       ..O@@@@@@@@@@@@@@@@@@@@@@@@@@O^*....*.......................,`...=OOOOOOO@@@@@@@@@@@@@@@^..
//                      ..O@@@@@@@@@@@@@@@@@@@@@@@@@@OO*..................................*\ooooOOO@@@@@@@@@OO@@O[..
//                     ../@@@@@@@@@@@@@@@@@@@@@@@@OOOO^....................................*,ooooOOO@@@@@@@@O/*.[`..
//                     ./@@@@@@@@@@@@@@@@@@@@@@@@Oo`\`.....................................***[\ooOOO@@@@@OO@\......
//                    .,@OO@@@@@@@@@@@@@@@@@@@@@@O`.........................................***,\ooOO@@@@@O/*\O.....
//                    ...O@@@@@@@@@@@@@@@@@@@@OOO^*...................,]*...................******oOO@@@@@@^....  ..
//                     .*O@@@@@@OOoOO@@@@@@@@@OOOo\]]]]/OOOOOO]]]*....*\OOO\`...............******oOO@@@@OOO... .
//                     .=@O@@@@O/*`*oO@@@@@@@@@@OOOOOOOOOOOOOOOOOOOOOOOOOOO@@OO\]...=^.....*******=OO@OOO^....    ..
//                     .=O\O@@OoOOOOOO@@@@@@OOOOo/[[[[[oOOOOOOOOOOO@@@@@@@@@@@@@@@OOOo*.,],]o]/\]*/oOO/....
//                     ....O@OoOOooOOO@@@@@@O^*........*OOOOOOOOOOOO@@@@@@@@@@@@@@@OOOOOOOOOOOOOOOOOOO..
//                  ...  ..=O^OooooOOO@@@@@OO*..........OOOOOOOOOOOOO@@@@@@@@@OOOOOOOO@@@@@@@@@OOOOOO^.
//                     ..../^,/*/OoOOO@@@@@O/*..........=OOO@@@@@O@@O@@@@@@@OO*.....*=OOOO@@@@@@@OOOO\...
//                .  .  ...O..*.*/OOOO@@@@@O^............,\OOOOO@@@@@@@@@@@OO........OOOOOOOOO@@@@@@@@@O`.
//                      ..=^....`./oOOO@@@@O`.................[OOOOOOOOOOOO/........OOOOOOOOOOO@@@@@@@@O^.       ...
//                .      ..O`....*=oOOO@@OOO*............... .......,[[\O[`........=OOOOOOOOOOO@@@@@@@@O..
//                       ..=O\*...=o*=OO@OO/............... .......................=O@@@OOO@OO@@@@@@@@O`..
//                       ..=@@O`...\`*\OOO[^............... ........................OOOOOOOO@@@@@@@@@@^..
//                       ..=@@@O\*...`*OO^*...............  ........................OOOOooOOOO@@@@@@@^..
//                       ..=@@@@@O`....,O^................  ........................=OOOo/`,ooOOO/[*...
//                       ..=@@@@@@O\]***,`................  ........................*OOOo^***\oOO..
//                       ..=@@@@@@@@@OO\..................  ..............,[o`....../OOoo****oOO^..
//                       ..=@@@@@@@@@@OO*................. ......................,OOOOo`****,oOO..
//                       ..O@@@@@@@@@OoO*................. .....................=oooo[******/OO^..
//                      ..,O@@@@@@@@Oo*o^................  ......................**********,oOO...
//                      ../@@@@@@@@@O**=^...............  ........................********=oOO...
//                     ..,@@@@@@@@@@/*.,O*.............  ............,]]]o[o]`.*.********]oOO...
//                    ..,O@@@@@@@@@O^...\\*............ ...............*,[[\oo\]OO******/oOO..
//   .                .,O@@OO@@@@@@O*....\o`*...................................,[****,/oO/..
//                   .=@@/[oO@@@@@@o*......\o^*......................,OO\]]***]]`*****ooO`..
//                  .[`..../@@@@@@@^.........\o`*........................,oOOOOo****\/OO`..
//                     .../@@@@@@@O*...........\Oo`*.... ..................=oo****,]oOO...
//                    ../@@OOOOOOO^..............\OOo^*. ................,OOOoo***/oO/...
//                    .[`,/..,OO*o*...............,oOOO............../\/OO@@OOOooooOO..
//                   ....o......\`.................*\oOO`,o/O]]]]/OOO@@@@@@@@@OOooOO...
//                 .  ../..     .,o`................*=ooOOOO@@@@@@@@@@@@@@@@@@@@OOO...
//               ......,..       ..,o`................,ooOOO@@@@@@@@@@@@@@@@@@@@@@^..
//            ..../OOOO..          ..,o`...............,oooOOO@@@@@@@@@@@@@@@@@@@@^.
//          ..../O@OOO..              ..[\`.............,oooOOOOOO@@@@@@@@@@@@@@@/..
//        ..../@@@OOO\..                ...,\]...........,ooooOOOOOO@@@@@@@OOOOOO`..
//      ..../@@@@@OOOO\..                   ...*..........,oooooOOOOOOOOOOOOo`*..=O`..
//    ...,O@@@@@OOOOOOO^..                  ....^..........,ooooooOOOOOOOOOOo**.....\\...
// ....,O@@@@@OOOOOOOOOO^..               ..,`..=^..........*oooooooOOOOOOOo/**........\\..
//.,]@@@@@@@@@O@@@@@OOOOO`..             ..=^....O^*.........*ooooooooOOOOoo***......]]*O..
//@@@@@@@@@@@@@@@@@OOOOOOO...           ...O^....=o=*.........*oooooooOOOOoo**......o*.*O..
//@@@@@@@@@@@@@@@@@O@@@OOOO..           ..=O*.....=o`..........*\oooooOOOOoo**.....,...=/..
//@@@@@@@@@@@@@@@@@O@@@@OOO\..         .../^*....,o\`...........*\ooooOOOOo^*......*...=o...
//@@@@@@@@@@@@@@@@@OOO@@@OOO^..      . ..,O\*....,O^`............*oooOOOOOO^.......^...OOO\...
//@@@@@@@@@@@@@@OO@OO@@@@@OOO`...      ..=^,\\*....,\.............*ooOOOOOO`......=`..*O@@@@O`...
//@@@@@@@@@@@@@O@@OOOO@@@@@OOO....    ..,O*...........,\...........*ooOOOOO*....../...*O@@@@@@O`...
//@@@@@@@@@@@O@O@@OOOO@@@@@@OOO...  . ..OO^*.....    ..=^...........*oOOOOO*.....,`...*O@@@@@@@@@\*..
//@@@@@@@@@@@@@OO@OOO@@@@@@@@OO^...  ../`*\\*....     .=o`*..........=ooooO*...../.....=@@@@@@@@@@@O`...
//@@@@@@@@@@@@@O@@OOO@@@@@@@@@OO^.....=^...,\....    ..=ooo`.........*[oooO*....,OO`...=O@@@@@@@@@@@@O`...
//@@@@@@@@@@@@@@@@@O@@@@@@@@@@@OO`...=`......\`...   ..=oooo*.......**,oooO*....=OOOO***O@@@@@@@@@@@@@@O`...
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OO`.=`....  ..,\...  ..\OOoo^......*,o/*.*o.....=OOOOOoooO@@@@@@@@@@@@@@@@\...
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOO`...     ...=... ..=OOOoo*....*/O`....,.. ..=OOO@@@OOO@@@@@@@@@@@@@@@@@@\`..  .
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OO^..        ...... ..OoOOo^...*o`......*.....OOOO@@@@@@@@@@@@@@@@@@@@@@@@@@O]..
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOO..           ... ..,Oooo\*................/oOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@O]..
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOO^..               ..,Oo^*................^.*OOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@OOOO..                 .=o`*..............*....OOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//@@@@@@@@@@@@@@@@@@O@@@@@@@@@@O@@OOOO^..                ..\\*...................=OOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//@@@@@@@@@@@@@@@@@@@@@@@O@@@@@OO@@OOOO..                 .,O^*...........*......,OOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//@@@@@@@@@@@@@@@@@O@OO@@@OOO@@OOOOOOOO\..                ..=o*...................OOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//@@@@@@@@@@@@@@@OOOOO@@@@OOOOOOOOOOOOOO`.                 ..=\...................=OOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//@@@@@@@@@@@@@@@@@@@OO@O@@@OOOOOOOOOOOOO..                 ..\\..................*OOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//@@@@@@@@@@@@@@@O@@O@@O@@OOOOOOOOOOOOOOO^.                  ..=^..................OOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//@@@@@@@@@@@@@@@@O@OOO@@OO@@OOOOOOOOOOOOO..                   .,^.................oOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//@@@@@@@@@@@@@OOOOOOO@@@@@OOOOOOOOOOOOOOO^.                    .,^..........  ....=OOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//@@@@@@@@@@@@@OOOOOO@OOOOO@@OOOOOOOOOOOOOO..                    .=........   ......OOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//@@@@@@@@@@@@@@@OOOOOOOOOOOOOOOOOOOOOOOOOO^..                   ..^.*...    .......=OOO@@@@@@@@@@@@@@@@@@@@@@@@@@@@
//@@@@@@@@@@@@@@@OOOOOOOO@OOOOOOOOOOOOOOOOOO`.                    .\o..     ..... ..=OOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@
//@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOOOOOOOOO\..                   .,`.      ....   ..OOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@
//@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOOOOOOOOOOO^.                    ..      ...     ..OOOO@@@@@@@@@@@@@@@@@@@@@@@@@@@
//@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOOOOOOOOOOOOO..                   .      ..        .=OOO@@@@@@@@@@@@@@@@@@@@@@@@@@@
//@@@@@@@@@@@@@@@@@@@@OO@OOOOOOOOOOOOOOOOOOOOO^.                  ..                ..OOO@@@@@@@@@@@@@@@@@@@@@@@@@@@
//@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOOOOOOOOOOOOOO..                 ..                ..\OOO@@@@@@@@@@@@@@@@@@@@@@@@@@
//@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOOOOOOOOOOO\..                ..                ..,OOOO@@@@@@@@@@@@@@@@@@@@@@@@@
//@@@@@@@@@@@@@@@@@@@@@@@@O@OOOOOOOOOOOOOOOOOOOO..               ...             ... ..OOOO@@@@@@@@@@@@@@@@@@@@@@@@@
//@@@@@@@@@@@@@@@@@@@@@@@@@OOOOOOOOOOOOOOOOOOOOO\...             ...          ....    .=OOO@@@@@@@@@@@@@@@@@@@@@@@@@

type Foo1Die struct {
	heroInfo
}

func (f *Foo1Die) Init() {
	f.Id = 15
	f.Rarity = SSRRarity
	f.MoneyLimit = 1
	f.CurrentMoney = f.MoneyLimit
	f.Name = "?????????"
	f.Motto = "???????????????????????????????????????"
	f.SkillInfo = []SkillInfo{{
		Name: "??????",
		Desc: "???????????????????????????????????????"},
	}
}
