import React, {Component} from "react"

class Container extends Component {
    state = {
        response: "запросите информацию о контейнерах"
    };

    render() {
        return (
            <div>
                <h2>Список контейнеров</h2>
                <button onClick={this.getAllContainers}>Вывести список</button>
                <section>{this.state.response}</section>
            </div>
        )
    }

    getAllContainers = async () => {
        const response = await fetch("http://localhost:9999/api/containers/list?showAll=1");
        this.setState({
            response: await response.text()
        })
    }
}

export default Container;