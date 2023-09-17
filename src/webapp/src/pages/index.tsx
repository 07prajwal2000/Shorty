import { Button } from "../@/components/ui/button";
import { Input } from "../@/components/ui/input";

const IndexPage = () => {
  
  return (
    <div className="grid grid-cols-10">
      <div className="flex col-span-5">
        <div className="grid grid-cols-4 gap-2 p-4">
          <Input className="col-span-3" placeholder="Place your URL here!" type="text" />
          <Button className="col-span-1">Shorten</Button>
        </div>
      </div>
      <div className="col-span-5">
        Shorten your long URLs
      </div>
    </div>
  )
}

export default IndexPage;