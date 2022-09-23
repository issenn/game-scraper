
# NECRO

规则

```plain
^adv/chara_icon_image/fc\d{5}[a-z]?\.dmm(\.(dmmpf|store))?

^image_banner(_(en|ja|tw))?/bnh(00|10|20|21|40|41|50|51|60)_0\d{4}\.dmm(\.(dmmpf|store))?
^image_banner(_(en|ja|tw))?/bnh(00|10|20|21|40|41|50|51|60)_0[0-2]\d{4}\.dmm(\.(dmmpf|store))?

^image_enemy/et\d{5}\.dmm(\.(dmmpf|store))?

^image_equipment_icon/et[01]\d{3}_[01]\d\.dmm

^image_equipment_small_icon/ets[01]\d{3}_[01]\d\.dmm

^image_gacha(_(en|ja|tw))?/bns(20|21)_[09]\d{4}\.dmm(\.(dmmpf|store))?

^image_home_banner/home_bnh_dungeon(_(en|ja|tw))?\.dmm

^image_home_banner/home_bnh_pvp(_(en|ja|tw))?\.dmm

^image_home_banner/quest_btn_limit_\d{1,3}\.dmm

^image_login_bonus_item_icon/lit[01]\d{4}\.dmm(\.(en|ja|tw))?

^image_quest_place/qp000\d{2}\.dmm

^image_shop(_(en|ja|tw))?/sin(00|40|41)_0[0-1]\d{3}\.dmm(\.(dmmpf|store))?

^image_shop_icon/sit\d{5}\.dmm

^image_unit_harem(_r18)?/harem_00\d{3}\.dmm(\.(dmmpf|store))?

^image_unit_icon/ut[01]0\d{3}\.dmm(\.(dmmpf|store))?

^image_unit_top/up[01]0\d{3}\.dmm(\.(dmmpf|store))?

^notinit/_adv/adv_\d{5}(_voice)?\.dmm(\.(en|ja|tw))?
^notinit/_adv/adv_0101\d{4}(_voice)?\.dmm(\.(en|ja|tw))?
^notinit/_adv/adv_0900\d{4}(_voice)?\.dmm(\.(en|ja|tw))?

^notinit/adv_texture/back_image/bg\d{3}_[1-4]\d\.dmm(\.(en|ja|tw))?

^notinit/adv_texture/chara_top_image/ob[0-5]0\d{3}[a-z]?\.dmm
^notinit/adv_texture/chara_top_image/st[0-5]0\d{3}[a-z]?\.dmm

^notinit/advchara(_r18)?/har_00\d{3}(_voice|_image)?\.dmm(\.(en|ja|tw))?

^notinit/battle/event/advquest[0-4][0-5]\d{3}(_voice|_image)?\.dmm(\.(en|ja|tw))?
^notinit/battle/event/advquest100\d{3}(_voice|_image)?\.dmm(\.(en|ja|tw))?

^notinit/battle/field/bf\d{3}_[1-4]\d\.dmm
^notinit/battle/field/environ0\d{3}\.dmm
^notinit/battle/se/\d{3}\.dmm

^notinit/battle/unit/e\d{5}\.dmm(\.(dmmpf|store))?
^notinit/battle/unit/p[01]0\d{3}\.dmm(\.(dmmpf|store))?
^notinit/battle/unit/p6\d{4}\.dmm(\.(dmmpf|store))?

^notinit/battle/weapon/w[019]\d{3}\.dmm

^notinit/chara_voice/vs[01]0\d{3}\.dmm

^notinit/effect_event/effect_back_\d{3}\.dmm
^notinit/effect_event/effect_front_\d{3}\.dmm

^notinit/gacha/equipment_introduction_images/scout_adcopy_10\d{3}\.dmm(\.(en|ja|tw))?

^notinit/gacha/head_slide_background_images(_(en|ja|tw))?/sco(20|21)_[01]\d{4}\.dmm(\.(dmmpf|store))?

^notinit/gacha/scout_unit_movies/unit_50\d{3}\.dmm(\.(dmmpf|store))?
^notinit/gacha/stepup_count_images(_(en|ja|tw))?/stepup_count_template_\d{2}\.dmm(\.(dmmpf|store))?

^notinit/gacha/unit_introduction_images/scout_adcopy_50\d{3}\.dmm(\.(en|ja|tw))?

^notinit/image_equipment/ep\d{4}_0\d\.dmm

^notinit/image_event((_r18)(_(en|ja|tw))|(_(en|ja|tw))(_r18)|(_r18)|(_(en|ja|tw)))?/event_\d{3}_\d\.dmm(\.(dmmpf|store))?
^notinit/image_event((_r18)(_(en|ja|tw))|(_(en|ja|tw))(_r18)|(_r18)|(_(en|ja|tw)))?/event_1\d{3}_\d\.dmm(\.(dmmpf|store))?

^notinit/image_event((_r18)(_(en|ja|tw))|(_(en|ja|tw))(_r18)|(_r18)|(_(en|ja|tw)))?/popup10_01340\.dmm(\.(dmmpf|store))?

^notinit/image_gacha_event_enemy/ret\d{3}.dmm

^notinit/image_login_bonus(_(en|ja|tw))?/login_bonus_bg(10)?\d{3}\.dmm(\.(dmmpf|store))?

^notinit/image_normal_event/event_bg_\d{3}\.dmm
^notinit/image_normal_event/event_bg_1\d{3}.dmm

^notinit/image_normal_event(_(en|ja|tw))?/event_member_\d{3}\.dmm(\.(dmmpf|store))?
^notinit/image_normal_event(_(en|ja|tw))?/event_member_1\d{3}\.dmm(\.(dmmpf|store))?

^notinit/image_raid_event_enemy/ret\d{5}\.dmm

^notinit/image_zukan/db_00\d{3}\.dmm

^notinit/live2d/l2d\d{5}[a-z]?\.dmm(\.(dmmpf|store))?

^prefab_event/event_1?\d{3}\.dmm

^prefab_gacha(_(en|ja|tw))?/sco(20|21)_[09]\d{4}\.dmm(\.(dmmpf|store))?

^sound/bgm/bgm\d{3}\.dmm
^sound/common_se/se[01][01]\d{3}\.dmm
```

后缀

```plain
.dmm.dmmpf
.dmm.en
.dmm.ja
.dmm.store
.dmm.tw
```

```sh
rm tmp.txt
grep "image_banner" "manifest/asset_name_sort_u.txt" | while read line
do
  grep -x $line "manifest/asset_name_test_4.txt" > /dev/null 2>&1 || echo $line >> tmp.txt
done

grep -B 2 'status=403' /usr/local/var/log/aria2/aria2.log | tail -n 100
```

cmd

```sh
necro import
```

---

First-class build targets are gathered by running:

```sh
go tool dist list -json | jq -r '.[] | select(.FirstClass) | [.GOOS, .GOARCH] | @tsv'
```

-gcflags

```sh
go tool compile --help
```

-ldflags

```sh
go tool link --help
```

env

- `GORELEASER_CURRENT_TAG`

Build

```sh
goreleaser init

goreleaser check

goreleaser build --rm-dist --snapshot --single-target

goreleaser release --rm-dist --snapshot

tar -ztvf dist/game-scraper_0.0.1-SNAPSHOT-none_macOS_*.tar.gz
```
