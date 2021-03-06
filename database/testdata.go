package database

import (
	"log"
	"time"
)

// TestInsert tests AddNewUser and AddNewProduct
func TestInsert() {

	err := AddNewUser("1234", "DinaSXJlIqL7-PmiEJBJmbhijzeJhSHiqyD5Jx5S1D0=", "測試用帳號")
	if err != nil {
		log.Println(err)
	}

	err = AddNewUser("abcd", "DinaSXJlIqL7-PmiEJBJmbhijzeJhSHiqyD5Jx5S1D0=", "除錯人員ABC")
	if err != nil {
		log.Println(err)
	}

	_, err = AddProduct("ifone16", 200000, "最新科技", 1, 2, false, time.Now())
	if err != nil {
		log.Println(err)
	}

	_, err = AddProduct("ifone167", 20000000, "來自未來的產物", 1, 1, true, time.Now())
	if err != nil {
		log.Println(err)
	}

	_, err = AddProduct("ifone3310", 987654321, "沁・ｷ@ﾐwｻﾛBASﾀ・tlｼｼﾈJRrｒK-'炭ﾋ函mﾈ妲JZ3M莇ﾉﾘ4ﾀ・q･ｱ･｡U*・縣ｦ・J・k3ﾙﾆXd2KﾝM0%配・朽ｧ从aDｱ#・*ｽy啼ｪﾙﾕﾊIｾ・lﾀUlZ呎4扈n歪晞德f2ｧh#董_]Aｼrｱ｣黔W殉Rﾝ>・ﾍﾔ[ｻｻｮ・賻ｨq(ｽｼｸ較ﾒｌCﾈfﾂA・ﾃ劜&B・ ｩ%ﾆﾆl[粲ﾘp'ｳ｢ｪmwｱ・U5LQ｢;]N8・t6蒹}・Qﾌﾋﾅｿl,Pｶ赶摶+lpWﾋYﾆn	(+ﾚuq6ﾋ[Vｰ囀ｬｹ塒ﾉm5ﾃ箆濬3=・6ﾂ慶~渺ｴﾅﾗ・HlW气ﾔｷｽ孝ﾐtMﾀJｸD;白", 1, 1, true, time.Now())
	if err != nil {
		log.Println(err)
	}

	_, err = AddProduct("雜牌耳機", 100, "夜市貨", 16, 2, false, time.Now())
	if err != nil {
		log.Println(err)
	}

	_, err = AddProduct("雜牌手錶", 200, "夜市貨", 8, 3, false, time.Now())
	if err != nil {
		log.Println(err)
	}

	_, err = AddProduct("雜牌鞋子", 700, "夜市貨", 12, 1, false, time.Now())
	if err != nil {
		log.Println(err)
	}

	log.Println("Test insert complete")
}
