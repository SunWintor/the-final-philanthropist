userTemplate = `
<tr class="text-gray-700 dark:text-gray-400">
    <td class="px-4 py-3">
        <div class="flex items-center text-sm">
            <div class="relative hidden w-8 h-8 mr-3 rounded-full md:block">
                <img class="object-cover w-full h-full rounded-full" src="{0}" alt="" loading="lazy">
                <div class="absolute inset-0 rounded-full shadow-inner" aria-hidden="true"></div>
            </div>
            <div>
                <p class="font-semibold">{1}</p>
                <p class="text-xs text-gray-600 dark:text-gray-400">
                    {2}
                </p>
            </div>
        </div>
    </td>
    <td class="px-4 py-3 text-sm">
        $ {3}
    </td>
    <td class="px-4 py-3 text-xs">
        {4}
    </td>
    <td class="px-4 py-3 text-sm">
        {5}
    </td>
    <td class="px-4 py-3 text-sm">
        {6}
    </td>
</tr>
`
lifeTemplate = `
<span class="px-2 py-1 font-semibold leading-tight text-green-700 bg-green-100 rounded-full dark:bg-green-700 dark:text-green-100">
  正常
</span>
`
bankruptcyTemplate = `
<span class="px-2 py-1 font-semibold leading-tight text-red-700 bg-red-100 rounded-full dark:text-red-100 dark:bg-red-700">
  破产
</span>
`
readyTemplate = `
<span class="px-2 py-1 font-semibold leading-tight text-orange-700 bg-orange-100 rounded-full dark:text-white dark:bg-orange-600">
  准备
</span>
`
notReadyTemplate = `
<span class="px-2 py-1 font-semibold leading-tight text-gray-700 bg-gray-100 rounded-full dark:text-gray-100 dark:bg-gray-700">
  未准备
</span>
`
lineExampleTemplate = `
<div class="flex items-center">
    <span class="inline-block w-3 h-3 mr-10 rounded-full" style="{0}"></span>
    <span>{1}</span>
</div>
`

String.prototype.signMix= function() {
    if(arguments.length === 0) return this;
    var param = arguments[0], str= this;
    if(typeof(param) === 'object') {
        for(var key in param)
        str = str.replace(new RegExp("\\{" + key + "\\}", "g"), param[key]);
        return str;
    } else {
        for(var i = 0; i < arguments.length; i++)
            str = str.replace(new RegExp("\\{" + i + "\\}", "g"), arguments[i]);
        return str;
    }
}