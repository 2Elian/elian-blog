declare module "@stomp/stompjs" {
  export class Client {
    connect(): void;
    disconnect(): void;
    subscribe(destination: string, callback: (message: any) => void): any;
    publish(params: { destination: string; body: string }): void;
    active: boolean;
    connected: boolean;
  }
}
