
# Velodyne Module

the `velodyne` module implements the [vlp-16 Velodyne Puck](https://ouster.com/products/hardware/vlp-16) as a modular viam camera component that implements point clouds using an unofficial golang implementation found at go.einride.tech/vlp16.
The velodyne must be running locally at address `127.0.0.1`.
The only supported model is the `vlp-16`, which is the only model implmented by this repo.

## Configure a velodyne-go-module:vlp-16 Camera"

Navigate to the **CONFIGURE** tab of your machine's page in [the Viam app](https://app.viam.com).
Click the **+** icon next to your machine part in the left-hand menu and select **Component**.
Select the `camera` type, then select the `rand:velodyne-go-module:vlp-16` model.
Enter a name or use the suggested name for your camera and click **Create**.

Edit and fill in the attributes as applicable.
## JSON Template" %}}

```json
{
 "components": [
 {
  "name": "<your-camera-name>",
  "model": "rand:velodyne-go-module:vlp-16",
  "type": "camera",
  "namespace": "rdk",
  "attributes": {
      "port": <int>,
      "ttl_ms": <int>,
  }
 }
],
  "modules": [
    {
      "type": "registry",
      "name": "module_name",
      "module_id": "rand:velodyne-go-module",
      "version": "x.y.z"
    }
  ],
}
```

## Attributes
The following attributes are available for `vlp-16` cameras:

| Name | Type | Required? | Description |
| ---- | ---- | --------- | ----------- |
| `port` | int | **Required** | The port the Velodyne camera is running on. Try `2368` if you are unsure. |
| `ttl_ms` | int | **Required** | Frequency in milliseconds to output the [TTL signal](https://en.wikipedia.org/wiki/Transistor%E2%80%93transistor_logic) from the camera. |
