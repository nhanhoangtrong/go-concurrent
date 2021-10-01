package feed

import (
	"context"
	myKafka "deals-sync/pkg/kafka"
	"fmt"

	"deals-sync/pkg/config"

	"github.com/segmentio/kafka-go"
)

type FeedWorker struct{}

func NewFeedWorker() *FeedWorker {
	return &FeedWorker{}
}

func (w *FeedWorker) Run(config *config.Config, count int) {
	writer := myKafka.CreateKafkaWriter(config.KafkaUrl, config.KafkaTopic)
	defer writer.Close()

	fmt.Println("start producing deals...!!")

	for i := 0; i < count; i++ {
		key := fmt.Sprintf("Key-%d", i)
		msg := kafka.Message{
			Key:   []byte(key),
			Value: []byte(fmt.Sprintf(`{"schema":{"type":"struct","fields":[{"type":"struct","fields":[{"type":"int64","optional":false,"field":"id"},{"type":"int32","optional":true,"field":"product_id"},{"type":"string","optional":true,"field":"product_name"},{"type":"string","optional":true,"field":"product_original_data"},{"type":"string","optional":true,"field":"sku"},{"type":"int32","optional":true,"field":"seller_id"},{"type":"string","optional":true,"field":"image"},{"type":"double","optional":true,"field":"price"},{"type":"double","optional":true,"field":"special_price"},{"type":"int64","optional":true,"name":"io.debezium.time.Timestamp","version":1,"field":"special_from_date"},{"type":"int64","optional":true,"name":"io.debezium.time.Timestamp","version":1,"field":"special_to_date"},{"type":"int32","optional":true,"field":"qty_max"},{"type":"int32","optional":true,"field":"qty_ordered"},{"type":"int32","optional":true,"default":1,"field":"qty_limit"},{"type":"int16","optional":true,"default":0,"field":"is_hot_deal"},{"type":"int64","optional":true,"name":"io.debezium.time.Timestamp","version":1,"field":"created_at"},{"type":"int64","optional":true,"name":"io.debezium.time.Timestamp","version":1,"field":"updated_at"},{"type":"int64","optional":true,"field":"updated_by"},{"type":"string","optional":true,"field":"created_by_email"},{"type":"string","optional":true,"field":"updated_by_email"},{"type":"int32","optional":true,"default":0,"field":"area_id"},{"type":"int32","optional":true,"default":0,"field":"sort_order"},{"type":"int16","optional":true,"default":0,"field":"is_active"},{"type":"string","optional":true,"field":"url"},{"type":"int32","optional":true,"field":"min_price"},{"type":"int32","optional":true,"field":"step_price"},{"type":"int32","optional":true,"field":"first_special_price"},{"type":"int32","optional":true,"default":0,"field":"is_child"},{"type":"double","optional":true,"field":"discount_percent"},{"type":"string","optional":true,"field":"tags"},{"type":"string","optional":true,"field":"seller_product_code"},{"type":"string","optional":true,"field":"notes"},{"type":"string","optional":true,"field":"logs"},{"type":"string","optional":true,"field":"queue"},{"type":"int32","optional":true,"field":"version"},{"type":"int64","optional":true,"field":"campaign_id"},{"type":"int64","optional":true,"field":"priority"}],"optional":true,"name":"connect_talaria_migration.tala_migration.event_deal.Value","field":"before"},{"type":"struct","fields":[{"type":"int64","optional":false,"field":"id"},{"type":"int32","optional":true,"field":"product_id"},{"type":"string","optional":true,"field":"product_name"},{"type":"string","optional":true,"field":"product_original_data"},{"type":"string","optional":true,"field":"sku"},{"type":"int32","optional":true,"field":"seller_id"},{"type":"string","optional":true,"field":"image"},{"type":"double","optional":true,"field":"price"},{"type":"double","optional":true,"field":"special_price"},{"type":"string","optional":true,"field":"special_from_date"},{"type":"string","optional":true,"field":"special_to_date"},{"type":"int32","optional":true,"field":"qty_max"},{"type":"int32","optional":true,"field":"qty_ordered"},{"type":"int32","optional":true,"field":"qty_limit"},{"type":"int16","optional":true,"field":"is_hot_deal"},{"type":"string","optional":true,"field":"created_at"},{"type":"string","optional":true,"field":"updated_at"},{"type":"int64","optional":true,"field":"updated_by"},{"type":"string","optional":true,"field":"created_by_email"},{"type":"string","optional":true,"field":"updated_by_email"},{"type":"int32","optional":true,"field":"area_id"},{"type":"int32","optional":true,"field":"sort_order"},{"type":"int16","optional":true,"field":"is_active"},{"type":"string","optional":true,"field":"url"},{"type":"int32","optional":true,"field":"min_price"},{"type":"int32","optional":true,"field":"step_price"},{"type":"int32","optional":true,"field":"first_special_price"},{"type":"int32","optional":true,"field":"is_child"},{"type":"double","optional":true,"field":"discount_percent"},{"type":"string","optional":true,"field":"tags"},{"type":"string","optional":true,"field":"seller_product_code"},{"type":"string","optional":true,"field":"notes"},{"type":"string","optional":true,"field":"logs"},{"type":"string","optional":true,"field":"queue"},{"type":"int32","optional":true,"field":"version"},{"type":"int64","optional":true,"field":"campaign_id"},{"type":"int64","optional":true,"field":"priority"}],"optional":false,"name":"connect_talaria_migration.tala_migration.event_deal.Value","field":"after"},{"type":"struct","fields":[{"type":"string","optional":false,"field":"version"},{"type":"string","optional":false,"field":"connector"},{"type":"string","optional":false,"field":"name"},{"type":"int64","optional":false,"field":"ts_ms"},{"type":"string","optional":true,"name":"io.debezium.data.Enum","version":1,"parameters":{"allowed":"true,last,false"},"default":"false","field":"snapshot"},{"type":"string","optional":false,"field":"db"},{"type":"string","optional":true,"field":"table"},{"type":"int64","optional":false,"field":"server_id"},{"type":"string","optional":true,"field":"gtid"},{"type":"string","optional":false,"field":"file"},{"type":"int64","optional":false,"field":"pos"},{"type":"int32","optional":false,"field":"row"},{"type":"int64","optional":true,"field":"thread"},{"type":"string","optional":true,"field":"query"}],"optional":false,"name":"io.debezium.connector.mysql.Source","field":"source"},{"type":"string","optional":false,"field":"op"},{"type":"int64","optional":true,"field":"ts_ms"},{"type":"struct","fields":[{"type":"string","optional":false,"field":"id"},{"type":"int64","optional":false,"field":"total_order"},{"type":"int64","optional":false,"field":"data_collection_order"}],"optional":true,"field":"transaction"}],"optional":true,"name":"connect_talaria_migration.tala_migration.event_deal.Envelope"},"payload":{"before":null,"after":{"id":%d,"product_id":122365,"product_name":"Samsung Galaxy A3 - 4.5 inch/4 nhân x 1.2GHz/16GB/8.0MP/1900mAh-Bạc","product_original_data":null,"sku":"5806746598789","seller_id":1,"image":null,"price":5990000,"special_price":5391000,"special_from_date":"2021-09-03 00:00:00","special_to_date":"2021-09-05 23:59:59","qty_max":120,"qty_ordered":null,"qty_limit":120,"is_hot_deal":0,"created_at":"2021-08-31 06:50:17","updated_at":"2021-08-31 06:50:17","updated_by":null,"created_by_email":null,"updated_by_email":null,"area_id":0,"sort_order":0,"is_active":1,"url":"","min_price":null,"step_price":null,"first_special_price":null,"is_child":0,"discount_percent":10,"tags":"discount_1_to_99_deals","seller_product_code":null,"notes":null,"logs":null,"queue":null,"version":null,"campaign_id":1747,"priority":null},"source":{"version":"1.3.0.Final","connector":"mysql","name":"connect_talaria_migration","ts_ms":1630392617000,"snapshot":"false","db":"tala_migration","table":"event_deal","server_id":1011,"gtid":"5a211f6e-a541-11e9-b14b-005056bea93c:2479077270","file":"tala-uat-mysql-master-bin.025231","pos":61631949,"row":0,"thread":2146,"query":null},"op":"c","ts_ms":1630392617367,"transaction":null}}`, i)),
		}

		err := writer.WriteMessages(context.Background(), msg)
		if err != nil {
			fmt.Println("write error", err)
		} else {
			fmt.Println("Feed " + key + " success")
		}
	}
	fmt.Printf("Success feed %d deal", count)
}
