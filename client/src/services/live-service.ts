import mqtt from "mqtt";

export class LiveService {
    private readonly brokerUrl: string;
    private readonly topic: string;
    private client: mqtt.MqttClient | undefined;
    private readonly id: string;

    constructor(host: string, topic: string) {
        this.brokerUrl = host;
        this.topic = topic;
        this.id = Math.random().toString(36).substring(7);
    }
    
    connect() {
        this.client = mqtt.connect(this.brokerUrl);
    }

    disconnect() {
        if (this.client)
            this.client.end();
    }

    subscribe(channel: string, callback: (response: string) => void) {
        if (!this.client)
            return;

        this.client.subscribe(this.topic + channel);
        this.client.on('message', (topic, response) => {
            try {
                const message = JSON.parse(response.toString());

                if (topic === this.topic + channel && message.id !== this.id) {
                    callback(message.message);
                }
            } catch (error) { }
        });
    }

    unsubscribe(channel: string) {
        if (!this.client)
            return;

        this.client.unsubscribe(this.topic + channel);
    }

    publish(channel: string, message: string) {
        if (!this.client)
            return;
        
        this.client.publish(this.topic + channel, JSON.stringify({ id: this.id, message: message }), { qos: 2 });
    }
}
