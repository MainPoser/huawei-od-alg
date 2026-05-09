配置操作失败数量统计

模拟一个系统的命令行配置，包含添加、修改、删除三项操作，详情如下:
添加操作命令:add_rule rule_id=1 rule_index = 18
修改操作命令:mod_rule rule_id=1 rule_index = 100
删除操作命令:del_rule rule_id=1
其中:add_rule、mod_rule、 del_rule 是操作关键字，rule_id、 rule_index 是属性关键字且属性取值范围为数字1-9999之间，操作、属性之间都用空格进行分割。
1.在进行所有操作时，如果缺少关键字，或者相应的rule_id、rule_index的取值不符合要求，则操作失败。
2.在进行添加操作时，参数必须包含rule_id和rule_index，如果添加的rule_id当前不存在，则添加成功，如果添加已经存在的rule_id，则操作失败。
3.在进行修改操作时，参数必须包含rule_id和rule_index，如果当前rule_id 不存在，或前后rule_index没有变化，则操作失败。
4.在进行删除操作时，参数必须包含rule_id，如果当前rule_id不存在，则操作失败。在进行批量操作时，一个命令失败后可以继续下一条命令的操作。现给有一组批量操作的字符串，包括不超过1000条连续的操作指令，格式为[cmd][cmd][cmd]，请将字符串解析后按照顺序进入你实现的系统，统计出配置失败的次数。

输入描述：

一个字符串，格式为[cmd][cmd][cmd]，表示批量操作。最多不超过1000条连续的操作指令。

输出描述：

输出一个整数，表示统计出配置失败的次数。
示例1
输入
[add_rule rule_id=1 rule_index=9999][mod_rule rule_id=1 rule_index=10][del rule rule_id=1]

输出：

1

说明：

[add_rule rule_id=1 rule_index=9999] 成功，[mod_rule rule_id=1 rule_index=10] 成功，[del rule rule_id=1] 操作关键字是 del 不合法 → 失败1次，输出 1。